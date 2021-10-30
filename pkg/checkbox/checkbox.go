package checkbox

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Checkbox struct {
	app.Compo
	base.JsUtil
	Id            string
	Label         string
	Checked       bool
	Disabled      bool
	Indeterminate bool
	Callback      func(input app.HTMLInput)
}

func (c *Checkbox) Render() app.UI {

	checkboxClass := "mdc-checkbox"
	if c.Disabled {
		checkboxClass = "mdc-checkbox mdc-checkbox--disabled"
	}

	input := app.Input().Type("checkbox").Class("mdc-checkbox__native-control").ID(c.Id + "-input").Disabled(c.Disabled)
	if c.Callback != nil {
		c.Callback(input)
	}
	if c.Indeterminate {
		input.DataSet("indeterminate", "true")
	}
	input.Checked(c.Checked)

	formField := app.Div().ID(c.Id + "-formField").Class("mdc-form-field")

	return formField.Body(
		app.Div().Class(checkboxClass).ID(c.Id).Body(
			input,
			app.Div().Class("mdc-checkbox__background").Body(
				app.Raw(SVG),
				app.Div().Class("mdc-checkbox__mixedmark"),
			),
			app.Div().Class("mdc-checkbox__ripple"),
		),
		app.If(c.Label == "").Else(
			app.Label().ID(c.Id+"-label").For(c.Id+"-input").Text(c.Label),
		),
	)
}

func (c *Checkbox) OnMount(ctx app.Context) {
	checkbox := c.JsNewAtPath("mdc.checkbox.MDCCheckbox", app.Window().GetElementByID(c.Id))
	formField := c.JsNewAtPath("mdc.formField.MDCFormField", app.Window().GetElementByID(c.Id+"-formField"))
	formField.Set("input", checkbox)
}

const SVG = `<svg class="mdc-checkbox__checkmark" viewBox="0 0 24 24">
<path class="mdc-checkbox__checkmark-path" fill="none" d="M1.73,12.91 8.1,19.28 22.79,4.59"></path>
</svg>`
