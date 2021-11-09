package slider

import (
	"fmt"
	"log"
	"strconv"

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

func (c *Continuous) OnMount(ctx app.Context) {
	c.api = c.JsNewAtPath(MDCSliderApi, c.JSValue())
	c.JSValue().Call("addEventListener", string(MDCSliderChange), app.FuncOf(c.event(ctx, MDCSliderChange)))
	ctx.Handle(string(MDCSliderValue), handleValue(c, c.api))
}

func (c *Continuous) ActionValue(ctx app.Context, value float64) {
	ctx.NewActionWithValue(string(MDCSliderChange), c,
		app.T("value", value),
		app.T("thumb", "2"),
	)
}

func (c *Continuous) event(ctx app.Context, change EventType) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		newChangeAction(ctx, c, args)
		return nil
	}
}

func newChangeAction(ctx app.Context, source interface{}, args []app.Value) {
	if len(args) > 0 && args[0].Get("detail").Truthy() {
		value := args[0].Get("detail").Get("value").Float()
		thumb := args[0].Get("detail").Get("value").Int()
		ctx.NewActionWithValue(string(MDCSliderChange), source,
			app.Tag{Name: "value", Value: fmt.Sprintf("%f", value)},
			app.Tag{Name: "thumb", Value: fmt.Sprintf("%d", thumb)},
		)
	}
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

func (cr *ContinuousRange) OnMount(ctx app.Context) {
	cr.api = cr.JsNewAtPath(MDCSliderApi, cr.JSValue())
	cr.JSValue().Call("addEventListener", string(MDCSliderChange), app.FuncOf(cr.event(ctx, MDCSliderChange)))
	ctx.Handle(string(MDCSliderValue), handleValue(cr, cr.api))
}

func (cr *ContinuousRange) ActionValue(ctx app.Context, thumb string, value float64) {
	ctx.NewActionWithValue(string(MDCSliderChange), cr,
		app.T("value", value),
		app.T("thumb", thumb),
	)
}

func (cr *ContinuousRange) event(ctx app.Context, change EventType) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		newChangeAction(ctx, cr, args)
		return nil
	}
}

func handleValue(compo interface{}, api app.Value) func(context app.Context, action app.Action) {
	return func(context app.Context, action app.Action) {
		if !api.Truthy() {
			log.Println("unable to handle event, no api set")
			return
		}
		if action.Name == string(MDCSliderValue) && action.Value == compo {
			value, err := strconv.ParseFloat(action.Tags.Get("value"), 64)
			if err != nil {
				log.Println(err)
				return
			}
			switch action.Tags.Get("thumb") {
			case "1":
				api.Call("setValueStart", value)
			case "2":
				api.Call("setValueEnd", value)
			}
		}
	}
}

type EventType string

const MDCSliderChange EventType = "MDCSlider:change"
const MDCSliderValue EventType = "MDCSlider:value"
