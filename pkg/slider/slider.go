package slider

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

const MDCSliderApi = "mdc.slider.MDCSlider"

type Continuous struct {
	app.Compo
	base.JsUtil
	Discrete bool
	Id       string
	Range    *InputRange
	api      app.Value
}

var _ app.Mounter = (*Continuous)(nil)
var _ app.Mounter = (*ContinuousRange)(nil)

func (c *Continuous) Render() app.UI {
	slider := app.Div().ID(c.Id).Class("mdc-slider")
	if c.Discrete {
		slider.Class("mdc-slider--discrete")
	}
	return slider.Body(
		c.Range,
		sliderTrack(),
		sliderThumb(c.Discrete, c.Range),
	)
}

func sliderThumb(discrete bool, r *InputRange) app.HTMLDiv {
	thumb := app.Div().Class("mdc-slider__thumb")
	thumbKnob := app.Div().Class("mdc-slider__thumb-knob")
	if discrete {
		return thumb.Body(
			app.Div().Class("mdc-slider__value-indicator-container").Aria("hidden", true).Body(
				app.Div().Class("mdc-slider__value-indicator").Body(
					app.Span().Class("mdc-slider__value-indicator-text").Text(r.Value),
				),
			),
			thumbKnob,
		)
	} else {
		return thumb.Body(thumbKnob)
	}
}

func sliderTrack() app.HTMLDiv {
	return app.Div().Class("mdc-slider__track").Body(
		app.Div().Class("mdc-slider__track--inactive"),
		app.Div().Class("mdc-slider__track--active").Body(
			app.Div().Class("mdc-slider__track--active_fill"),
		),
	)
}

func (c *Continuous) OnMount(_ app.Context) {
	c.api = c.JsNewAtPath(MDCSliderApi, c.JSValue())
}

type ContinuousRange struct {
	app.Compo
	base.JsUtil
	Id       string
	Discrete bool
	RangeOne *InputRange
	RangeTwo *InputRange
	api      app.Value
}

func (cr *ContinuousRange) Render() app.UI {
	slider := app.Div().ID(cr.Id).Class("mdc-slider mdc-slider--range")
	if cr.Discrete {
		slider.Class("mdc-slider--discrete")
	}
	return slider.Body(
		cr.RangeOne, cr.RangeTwo,
		sliderTrack(),
		sliderThumb(cr.Discrete, cr.RangeOne), sliderThumb(cr.Discrete, cr.RangeTwo),
	)
}

func (cr *ContinuousRange) OnMount(_ app.Context) {
	cr.api = cr.JsNewAtPath(MDCSliderApi, cr.JSValue())
}
