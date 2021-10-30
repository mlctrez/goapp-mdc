package main

import (
	"crypto/rand"
	"fmt"
	"image"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/mlctrez/goapp-mdc/pkg/banner"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/card"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
	"github.com/mlctrez/goapp-mdc/pkg/dialog"
	"github.com/mlctrez/goapp-mdc/pkg/example"
	"github.com/mlctrez/goapp-mdc/pkg/fab"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/tab"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &index{})
	app.Route("/demo", &Demo{})
	app.RunWhenOnBrowser()
	httpServer()
}

type pulseImage struct {
	app.Compo
	base.JsUtil
	id      string
	opacity float32
	ripple  app.Value
}

func (p *pulseImage) Render() app.UI {
	return app.Div().ID(p.id).Class("mdc-ripple-surface").Body(
		app.Img().Src("/web/soon.png").Style("opacity", fmt.Sprintf("%f", p.opacity)),
	)
}

func (p *pulseImage) getBoundingRect() image.Rectangle {
	cr := p.JSValue().Call("getBoundingClientRect")
	left := cr.Get("left").Int()
	right := cr.Get("right").Int()
	top := cr.Get("top").Int()
	bottom := cr.Get("bottom").Int()
	return image.Rect(left, top, right, bottom)
}

// TODO: is using crypto/rand in wasm ok size and performance wise?
func randomPointIn(rect image.Rectangle) image.Point {
	center := image.Pt(rect.Min.X+rect.Dx()/2, rect.Min.Y+rect.Dy()/2)
	rx, err := rand.Int(rand.Reader, big.NewInt(int64(rect.Dx())))
	if err != nil {
		return center
	}
	ry, err := rand.Int(rand.Reader, big.NewInt(int64(rect.Dx())))
	if err != nil {
		return center
	}
	return image.Pt(rect.Min.X+int(rx.Int64()), rect.Min.Y+int(ry.Int64()))
}

func (p *pulseImage) OnMount(ctx app.Context) {
	p.ripple = p.JsNewAtPath("mdc.ripple.MDCRipple", p.JSValue())

	if p.ripple.Truthy() {
		ctx.Async(func() {
			// wait a bit for page load
			time.Sleep(1 * time.Second)

			// fade in
			for p.opacity = 0; p.opacity < 1; p.opacity += 0.025 {
				p.Update()
				time.Sleep(20 * time.Millisecond)
			}

			rect := p.getBoundingRect()
			if rect.Dx() < 500 {
				// skip if image has not loaded in 2 seconds
				return
			}

			// touch a random point
			client := randomPointIn(rect)
			p.JSValue().Call("dispatchEvent", p.createEvent("mousedown", client.X, client.Y))
			time.Sleep(100 * time.Millisecond)
			p.JSValue().Call("dispatchEvent", p.createEvent("mouseup", client.X, client.Y))

		})
	}
}

func (p *pulseImage) createEvent(name string, x int, y int) app.Value {
	return p.JsNewAtPath("MouseEvent", name, map[string]interface{}{
		"bubbles":    true,
		"cancelable": true,
		"clientX":    x,
		"clientY":    y,
		"button":     0,
	})
}

type index struct {
	app.Compo
}

func (i *index) Render() app.UI {
	return app.Div().Class("goapp-mdc-coming-soon").Body(
		app.Div().Class("coming-soon-buttons").Body(
			&button.Button{Id: "btn1", Label: "Coming", Outlined: true},
			&button.Button{Id: "btn2", Label: "Soon", Raised: true, Callback: func(button app.HTMLButton) {
				button.OnClick(func(ctx app.Context, e app.Event) {
					ctx.Navigate("/demo")
				})
			}},
		),
		&pulseImage{},
	)
}

func (i *index) OnPreRender(ctx app.Context) {
	i.initPage(ctx)
}

func (i *index) OnNav(ctx app.Context) {
	i.initPage(ctx)
}

func (i *index) initPage(ctx app.Context) {
	page := ctx.Page()
	page.SetAuthor("mlctrez")
	page.SetKeywords("go, golang, go-app, pwa, wasm, material design components")
	page.SetTitle("goapp-mdc")
	page.SetDescription("Material Design Components for go-app")
}

func buildDemoMap() map[string]app.UI {
	return map[string]app.UI{
		"banner":   &banner.Demo{},
		"button":   &button.Demo{},
		"card":     &card.Demo{},
		"icon":     &icon.Demo{},
		"checkbox": &checkbox.Demo{},
		"dialog":   &dialog.Demo{},
		"form":     &example.Example{},
		"fab":      &fab.Demo{},
	}
}

type Demo struct {
	app.Compo
	bar         *tab.Bar
	ActiveIndex int
	DemoMap     map[string]app.UI
}

func (d *Demo) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (d *Demo) sortedDemoNames() []string {
	var sortedNames []string
	for k := range d.DemoMap {
		sortedNames = append(sortedNames, k)
	}
	sort.Strings(sortedNames)
	return sortedNames
}

func (d *Demo) Render() app.UI {

	once := sync.Once{}
	once.Do(func() {
		d.DemoMap = buildDemoMap()
		var tabs []*tab.Tab
		var index int
		for _, groupName := range d.sortedDemoNames() {
			newTab := tab.NewTab(groupName, index)
			if d.ActiveIndex == index {
				newTab.Active()
			}
			tabs = append(tabs, newTab)
			index++
		}
		d.bar = tab.NewBar("demoPageBar", tabs)
		d.bar.ActivateCallback(func(index int) {
			d.ActiveIndex = index
			d.Update()
		})
	})

	var content []app.UI
	content = append(content, d.bar)
	content = append(content, d.DemoMap[d.sortedDemoNames()[d.ActiveIndex]])

	return app.Div().Body(content...)
}
