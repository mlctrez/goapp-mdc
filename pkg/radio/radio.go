package radio

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/formfield"
)

type Radio struct {
	app.Compo
	autoinit.AutoInit
	Id       string
	Name     string
	Value    string
	Checked  bool
	Disabled bool
	api      app.Value
}

func (r *Radio) Render() app.UI {
	mdcRadio := app.Div()
	mdcRadio.Class("mdc-radio")
	if r.Disabled {
		mdcRadio.Class("mdc-radio--disabled")
	}

	Api.DataMdcAutoInitDiv(mdcRadio)

	radio := app.Input().Type("radio")
	radio.Class("mdc-radio__native-control")
	radio.ID(r.Id)
	radio.Name(r.Name)
	radio.Value(r.Value)
	if r.Checked {
		radio.Checked(r.Checked)
	}

	mdcRadio.Body(radio, Background(), Ripple())
	return mdcRadio
}

var _ formfield.Component = (*Radio)(nil)

func (r *Radio) OnMount(_ app.Context) {
	r.api = r.AutoInitComponent(r.JSValue(), Api)
	// required since value is only mapped to the js dom value not the html
	// default value by go-app.
	// mdc.slider.MDCSlider requires that the value="x" be present also
	r.JSValue().Call("setAttribute", "value", r.Value)
}

func Background() app.UI {
	return app.Div().Class("mdc-radio__background").Body(
		app.Div().Class("mdc-radio__outer-circle"),
		app.Div().Class("mdc-radio__inner-circle"),
	)
}

func Ripple() app.UI {
	return app.Div().Class("mdc-radio__ripple")
}

const Api = autoinit.MDCRadio

func (r *Radio) MdcAPI() app.Value {
	return r.api
}

func (r *Radio) LabelID() string {
	return r.Id
}
