package icon

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

// Toggle represents a mdc-icon-button with toggle states
type Toggle struct {
	app.Compo
	base.JsUtil
	autoinit.AutoInit
	Id      string
	Icon    MaterialIcon
	IconOff MaterialIcon
	State   bool
	AriaOn  string
	AriaOff string
	api     app.Value
}

func (b *Toggle) Render() app.UI {
	button := app.Button().Class("mdc-icon-button")
	ToggleApi.DataMdcAutoInitButton(button)
	if b.Id != "" {
		button.ID(b.Id)
	}

	button.DataSet("aria-label-on", b.AriaOn)
	button.DataSet("aria-label-off", b.AriaOff)
	if b.State {
		button.Class("mdc-icon-button--on")
		button.Aria("label", b.AriaOn)
	} else {
		button.Aria("label", b.AriaOff)
	}

	button.Body(
		RippleDiv(),
		b.Icon.I().Class("mdc-icon-button__icon mdc-icon-button__icon--on"),
		b.IconOff.I().Class("mdc-icon-button__icon"),
	)

	return button
}

var _ app.Mounter = (*Toggle)(nil)

const ToggleApi = autoinit.MDCIconButtonToggle

type Event string

const ToggleChange Event = "MDCIconButtonToggle:change"
const ToggleValue Event = "MDCIconButtonToggle:value"

func (b *Toggle) OnMount(ctx app.Context) {
	b.api = b.AutoInitComponent(b.JSValue(), ToggleApi)
	b.JSValue().Call("addEventListener", string(ToggleChange), app.FuncOf(b.event(ctx, ToggleChange)))
	ctx.Handle(string(ToggleValue), b.handleValueChange)
}

func (b *Toggle) event(ctx app.Context, change Event) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		isOnJs := b.JsValueAt(args[0], "detail.isOn", true)
		if isOnJs.IsUndefined() {
			return nil
		}
		b.State = isOnJs.Bool()
		state := "off"
		if b.State == true {
			state = "on"
		}
		ctx.NewActionWithValue(string(ToggleChange), b, app.T("state", state))
		return nil
	}
}

func (b *Toggle) handleValueChange(ctx app.Context, action app.Action) {
	if action.Name == string(ToggleValue) && action.Value == b {
		if b.api == nil {
			log.Println("unable to change toggle state, api is nil")
		}
		state := action.Tags.Get("state")
		switch state {
		case "on":
			b.api.Set("on", true)
		case "off":
			b.api.Set("on", false)
		}
	}
}
