package icon

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

// Button represents a material icon-button.
type Button struct {
	app.Compo
	base.JsUtil
	// Id is the id of the button which is required for a toggle button for binding MDCIconButtonToggle.
	// When present the ripple effect is also bound to the button.
	Id string
	// Icon represents the material icon used in the icon button state or the on state of a toggle button
	Icon MaterialIcon
	// IconOff represents the icon used in the off state of a toggle button
	IconOff MaterialIcon
	// State represents the current button state, set to true for rendering a button toggle in the initial state of on
	State bool
	// AriaOn represents the aria-label for the on->off action. i.e. "Remove from favorites"  used only in toggle button
	AriaOn string
	// AriaOff represents the aria-label for the off->on action. i.e. "Add to favorites"  used in toggle button and normal button
	AriaOff string
	// jsTarget is the new MDCIconButtonToggle on the root button
	jsTarget app.Value
	// Callback allows access to the button events when in non toggle mode
	Callback func(button app.HTMLButton)
	// ButtonToggleChange is called when the MDCIconButtonToggle:change event occurs on the toggle button
	ButtonToggleChange func(isOn bool)
}

func (b *Button) Render() app.UI {

	button := app.Button().Class("mdc-icon-button")
	if b.Id != "" {
		button.ID(b.Id)
	}
	if b.IconOff == "" {
		button.Class("material-icons")
		button.Aria("label", b.AriaOff)
		button.Body(
			app.Div().Class("mdc-icon-button__ripple"),
			app.Text(b.Icon),
		)
	} else {
		button.DataSet("aria-label-on", b.AriaOn)
		button.DataSet("aria-label-off", b.AriaOff)
		if b.State {
			button.Class("mdc-icon-button--on")
			button.Aria("label", b.AriaOn)
		} else {
			button.Aria("label", b.AriaOff)
		}

		button.Body(
			app.Div().Class("mdc-icon-button__ripple"),
			b.Icon.I().Class("mdc-icon-button__icon mdc-icon-button__icon--on"),
			b.IconOff.I().Class("mdc-icon-button__icon"),
		)

	}
	if b.Callback != nil {
		b.Callback(button)
	}
	return button
}

func (b *Button) OnMount(_ app.Context) {
	if b.IconOff != "" {

		element := b.JSValue()
		b.jsTarget = b.JsNewAtPath("mdc.iconButton.MDCIconButtonToggle", element)
		element.Call("addEventListener", "MDCIconButtonToggle:change", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			b.State = args[0].Get("detail").Get("isOn").Bool()
			b.ButtonToggleChange(b.State)
			return nil
		}))
	}
	if b.Id != "" {
		ripple := b.MDCRipple(b.Id)
		if ripple.Truthy() {
			ripple.Set("unbounded", true)
		}
	}
}

func (b *Button) SetState(isOn bool) {
	b.State = isOn
	if b.jsTarget.Truthy() {
		b.jsTarget.Set("on", isOn)
	}
}
