package checkbox

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/formfield"
)

type Checkbox struct {
	app.Compo
	base.JsUtil
	autoinit.AutoInit
	Id            string
	Checked       bool
	Disabled      bool
	Indeterminate bool
	Callback      func(input app.HTMLInput)
	api           app.Value
}

func (c *Checkbox) Render() app.UI {
	mdcCheckBox := app.Div().Class("mdc-checkbox")
	Api.DataMdcAutoInitDiv(mdcCheckBox)
	if c.Disabled {
		mdcCheckBox.Class("mdc-checkbox mdc-checkbox--disabled")
	}
	input := c.buildInput()

	mdcCheckBox.Body(
		input,
		Background(),
		app.Div().Class("mdc-checkbox__ripple"),
	)
	return mdcCheckBox
}

func (c *Checkbox) buildInput() app.HTMLInput {
	input := app.Input().Type("checkbox")
	input.Class("mdc-checkbox__native-control")
	input.ID(c.Id)
	input.Checked(c.Checked)
	if c.Disabled {
		input.Disabled(c.Disabled)
	}
	if c.Indeterminate {
		input.DataSet("indeterminate", "true")
	}
	if c.Callback != nil {
		c.Callback(input)
	}
	return input
}

func (c *Checkbox) OnMount(ctx app.Context) {
	c.api = c.AutoInitComponent(c.JSValue(), Api)
}

func Background() app.HTMLDiv {
	return app.Div().Class("mdc-checkbox__background").Body(
		app.Raw(SVG),
		app.Div().Class("mdc-checkbox__mixedmark"),
	)
}

const SVG = `<svg class="mdc-checkbox__checkmark" viewBox="0 0 24 24">` +
	`<path class="mdc-checkbox__checkmark-path" fill="none" d="M1.73,12.91 8.1,19.28 22.79,4.59"></path></svg>`

var _ formfield.Component = (*Checkbox)(nil)
var Api = autoinit.MDCCheckbox
func (c *Checkbox) LabelID() string   { return c.Id }
func (c *Checkbox) MdcAPI() app.Value { return c.api }
