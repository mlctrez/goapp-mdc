package switchm

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type MDCSwitch struct {
	app.Compo
	base.JsUtil
	autoinit.AutoInit
	Selected bool
	Disabled bool
	api      app.Value
}

func (s *MDCSwitch) Render() app.UI {
	button := app.Button().Class("mdc-switch")
	button.Type("button").Role("switch")
	button.DataSet("mdc-auto-init", string(autoinit.MDCSwitch))
	if s.Disabled {
		button.Disabled(s.Disabled)
	}
	if s.Selected {
		button.Aria("checked", "true")
		button.Class("mdc-switch--selected")
	} else {
		button.Aria("checked", "false")
		button.Class("mdc-switch--unselected")
	}

	return addBody(button)
}

func addBody(button app.HTMLButton) app.HTMLButton {
	return button.Body(
		app.Div().Class("mdc-switch__track"),
		app.Div().Class("mdc-switch__handle-track").Body(
			app.Div().Class("mdc-switch__handle").Body(
				app.Div().Class("mdc-switch__shadow").Body(
					app.Div().Class("mdc-elevation-overlay"),
				),
				app.Div().Class("mdc-switch__ripple"),
				app.Div().Class("mdc-switch__icons").Body(
					app.Raw(SwitchOnSvg),
					app.Raw(SwitchOffSvg),
				),
			),
		),
	)
}

var _ app.Mounter = (*MDCSwitch)(nil)

func (s MDCSwitch) OnMount(ctx app.Context) {
	s.api = s.AutoInitComponent(s.JSValue(), autoinit.MDCSwitch)
}

const SwitchOnSvg = `<svg class="mdc-switch__icon mdc-switch__icon--on" viewBox="0 0 24 24">` +
	`<path d="M19.69,5.23L8.96,15.96l-4.23-4.23L2.96,13.5l6,6L21.46,7L19.69,5.23z"/></svg>`

const SwitchOffSvg = `<svg class="mdc-switch__icon mdc-switch__icon--off" viewBox="0 0 24 24">` +
	`<path d="M20 13H4v-2h16v2z" /></svg>`
