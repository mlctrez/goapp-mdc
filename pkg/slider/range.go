package slider

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type InputRange struct {
	app.Compo
	base.JsUtil
	Id    string
	Name  string
	Label string
	Min   float64
	Max   float64
	Value float64
	Step  float64
}

var _ app.Mounter = (*InputRange)(nil)

func (r *InputRange) Render() app.UI {
	input := app.Input().Type("range").Class("mdc-slider__input")
	input.Name(r.Name).Aria("label", r.Label)
	if r.Id != "" {
		input.ID(r.Id)
	}
	input.Min(r.Min).Max(r.Max).Value(r.Value).Step(r.Step)

	return input
}

func (r *InputRange) OnMount(context app.Context) {
	// required since value is only mapped to the js dom value not the html
	// default value by go-app.
	// mdc.slider.MDCSlider requires that the value="x" be present also
	r.JSValue().Call("setAttribute", "value", r.Value)
}
