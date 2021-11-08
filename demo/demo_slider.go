package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/slider"
)

type SliderDemo struct {
	app.Compo
	base.JsUtil
}

func (d *SliderDemo) Render() app.UI {

	defaults := func(r *slider.InputRange) *slider.InputRange { r.Min = 0; r.Max = 100; r.Step = 5; return r }

	label := "Continuous slider demo"
	continuous := &slider.Continuous{Id: "C1",
		Range: defaults(&slider.InputRange{Id: "C1R", Name: "volume", Label: label, Value: 25}),
	}

	labelDiscrete := "Discrete slider demo"
	discrete := &slider.Continuous{Id: "D1", Discrete: true,
		Range: defaults(&slider.InputRange{Id: "D1R", Name: "volume", Label: labelDiscrete, Value: 75}),
	}

	labelRange := "Continuous range slider demo"
	continuousRange := &slider.ContinuousRange{Id: "CR1",
		RangeOne: defaults(&slider.InputRange{Id: "CR1R0", Name: "start", Label: labelRange, Value: 15}),
		RangeTwo: defaults(&slider.InputRange{Id: "CR1R1", Name: "end", Label: labelRange, Value: 85}),
	}

	labelDiscreteRange := "Discrete range slider demo"
	discreteRange := &slider.ContinuousRange{Id: "CR1", Discrete: true,
		RangeOne: defaults(&slider.InputRange{Id: "CR1R0", Name: "start", Label: labelDiscreteRange, Value: 35}),
		RangeTwo: defaults(&slider.InputRange{Id: "CR1R1", Name: "end", Label: labelDiscreteRange, Value: 65}),
	}

	body := layout.Grid().Body(
		layout.Inner().Body(
			layout.CellWide().Body(continuous, app.Text("Continuous slider")),
			layout.CellWide().Body(discrete, app.Text("Discrete slider")),
			layout.CellWide().Body(continuousRange, app.Text("Continuous range slider")),
			layout.CellWide().Body(discreteRange, app.Text("Discrete range slider")),
		),
	)

	return PageBody(body)

}
