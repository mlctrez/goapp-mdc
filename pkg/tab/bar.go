package tab

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Bar struct {
	app.Compo
	base.JsUtil
	id       string
	tabs     []*Tab
	jsApi    app.Value
	activate func(indx int)
}

func NewBar(id string, tabs []*Tab) *Bar {
	return &Bar{id: id, tabs: tabs}
}

func (b *Bar) ActivateCallback(f func(index int)) *Bar {
	b.activate = f
	return b
}

func (b *Bar) Render() app.UI {
	scrollContent := app.Div().Class("mdc-tab-scroller__scroll-content")
	var body []app.UI
	for _, tab := range b.tabs {
		body = append(body, tab)
	}
	scrollContent.Body(body...)
	return app.Div().ID(b.id).Class("mdc-tab-bar").Attr("role", "tablist").Body(
		app.Div().Class("mdc-tab-scroller").Body(
			app.Div().Class("mdc-tab-scroller__scroll-area").Body(
				scrollContent,
			),
		))
}

func (b *Bar) activateTab(tabIndex int) {
	b.jsApi.Call("activateTab", tabIndex)
}

func (b *Bar) OnMount(ctx app.Context) {
	element := app.Window().GetElementByID(b.id)
	b.jsApi = b.JsNewAtPath("mdc.tabBar.MDCTabBar", element)
	element.Call("addEventListener", "MDCTabBar:activated", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		if b.activate != nil {
			b.activate(args[0].Get("detail").Get("index").Int())
		}
		return nil
	}))
}
