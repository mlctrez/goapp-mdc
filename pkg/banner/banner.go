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

	content := app.Div().Class("mdc-banner__content").
		Attr("role", "alertdialog").Aria("live", "assertlive")

	textWrapper := app.Div().Class("mdc-banner__graphic-text-wrapper").Body(
		app.Div().Class("mdc-banner__text").Text(b.Text),
	)
	actions := app.Div().Class("mdc-banner__actions").Body(b.Buttons...)

	content.Body(textWrapper, actions)

	if b.Fixed {
		content = app.Div().Class("mdc-banner__fixed").Body(content)
	}

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

func (b *Banner) OnMount(ctx app.Context) {
	e := b.JSValue()
	b.jsApi = b.JsNewAtPath("mdc.banner.MDCBanner", e)
	e.Call("addEventListener", string(Opening), app.FuncOf(b.event(ctx, Opening)))
	e.Call("addEventListener", string(Opened), app.FuncOf(b.event(ctx, Opened)))
	e.Call("addEventListener", string(Closing), app.FuncOf(b.event(ctx, Closing)))
	e.Call("addEventListener", string(Closed), app.FuncOf(b.event(ctx, Closed)))
}

type EventType string

const Opening EventType = "MDCBanner:opening"
const Opened EventType = "MDCBanner:opened"
const Closing EventType = "MDCBanner:closing"
const Closed EventType = "MDCBanner:closed"

func (b *Banner) event(ctx app.Context, event EventType) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		// determine what button was clicked if any from the event args
		reason := -1
		if len(args) > 0 {
			detailReason := b.JsValueAt(args[0], "detail.reason", false)
			if !detailReason.IsUndefined() {
				reason = detailReason.Int()
			}
		}
		ctx.NewActionWithValue(string(event), b, app.T("reason", reason))
		return nil
	}
}
