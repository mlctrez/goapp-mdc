package banner

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Banner struct {
	app.Compo
	base.JsUtil
	Id       string
	Text     string
	Centered bool
	Fixed    bool
	Buttons  []app.UI
	jsApi    app.Value
	OnClosed func(string)
}

func (b *Banner) Render() app.UI {

	wrapper := app.Div().Class("mdc-banner__graphic-text-wrapper").Body(
		app.Div().Class("mdc-banner__text").Text(b.Text),
	)
	actions := app.Div().Class("mdc-banner__actions").Body(b.Buttons...)

	content := app.Div().Class("mdc-banner__content").Attr("role", "alertdialog").Aria("live", "assertlive")
	content.Body(wrapper, actions)

	banner := app.Div().Class("mdc-banner").ID(b.Id).Body(content)
	if b.Centered {
		banner.Class("mdc-banner--centered")
	}

	return banner
}

func (b *Banner) Open() {
	if b.jsApi.Truthy() {
		b.jsApi.Call("open")
	}
}

func (b *Banner) Close() {
	if b.jsApi.Truthy() {
		b.jsApi.Call("close")
	}
}

func (b *Banner) OnMount(_ app.Context) {
	element := app.Window().GetElementByID(b.Id)
	b.jsApi = b.JsNewAtPath("mdc.banner.MDCBanner", element)

	element.Call("addEventListener", "MDCBanner:closed", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		if b.OnClosed != nil {
			closeReason := args[0].Get("detail").Get("reason")
			closeReasonInt := -1
			if !closeReason.IsUndefined() {
				closeReasonInt = closeReason.Int()
			}
			switch closeReasonInt {
			case 0:
				b.OnClosed("primary")
			case 1:
				b.OnClosed("secondary")
			default:
				b.OnClosed("undefined")
			}
		}
		return nil
	}))

}

type Fixed struct {
	Banner
}

func (f *Fixed) Render() app.UI {

	wrapper := app.Div().Class("mdc-banner__graphic-text-wrapper").Body(
		app.Div().Class("mdc-banner__text").Text(f.Text),
	)
	actions := app.Div().Class("mdc-banner__actions").Body(f.Buttons...)

	content := app.Div().Class("mdc-banner__content").Attr("role", "alertdialog").Aria("live", "assertlive")
	content.Body(wrapper, actions)

	content = app.Div().Class("mdc-banner__fixed").Body(content)

	banner := app.Div().Class("mdc-banner").ID(f.Id).Body(content)
	if f.Centered {
		banner.Class("mdc-banner--centered")
	}

	return banner
}
