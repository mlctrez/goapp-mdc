package older

import (
	"crypto/rand"
	"fmt"
	"image"
	"math/big"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type PulseImage struct {
	app.Compo
	base.JsUtil
	id      string
	opacity float32
	ripple  app.Value
}

func (p *PulseImage) Render() app.UI {
	return app.Div().ID(p.id).Class("mdc-ripple-surface").Body(
		app.Img().Src("/web/soon.png").Style("opacity", fmt.Sprintf("%f", p.opacity)),
	)
}

func (p *PulseImage) getBoundingRect() image.Rectangle {
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

func (p *PulseImage) OnMount(ctx app.Context) {
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

func (p *PulseImage) createEvent(name string, x int, y int) app.Value {
	return p.JsNewAtPath("MouseEvent", name, map[string]interface{}{
		"bubbles":    true,
		"cancelable": true,
		"clientX":    x,
		"clientY":    y,
		"button":     0,
	})
}
