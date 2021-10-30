package progress

import (
	"fmt"
	"math"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Circular struct {
	app.Compo
	base.JsUtil
	id       string
	size     int
	label    string
	progress float64
	onmount  func(ctx app.Context)
	colors   [4]string
	target   app.Value
}

func NewCircular(id string, size int) *Circular {
	return &Circular{id: id, size: size, target: app.Undefined()}
}

func (c *Circular) Label(label string) *Circular {
	c.label = label
	return c
}

func (c *Circular) Colors(colors [4]string) *Circular {
	c.colors = colors
	return c
}

func (c *Circular) WithOnMount(cb func(ctx app.Context)) *Circular {
	c.onmount = cb
	return c
}

func (c *Circular) Render() app.UI {

	root := app.Div().ID(c.id)
	root.Class("mdc-circular-progress", "mdc-circular-progress--closed").Aria("hidden", "true")
	sizePx := fmt.Sprintf("%dpx", c.size)
	root.Style("width", sizePx).Style("height", sizePx)
	root.Aria("valuemin", "0").Aria("valuemax", "1").Aria("valuenow", "0")
	if c.label != "" {
		root.Aria("label", c.label)
	}

	determinateContainer := app.Div().Class("mdc-circular-progress__determinate-container").Body(app.Raw(c.DeterminateSVG()))
	indeterminateContainer := app.Div().Class("mdc-circular-progress__indeterminate-container")
	if c.colors[0] == "" {
		// single color spinner
		indeterminateContainer.Body(
			app.Div().Class("mdc-circular-progress__spinner-layer").Body(
				app.Div().Class("mdc-circular-progress__circle-clipper", "mdc-circular-progress__circle-left").Body(app.Raw(c.IndeterminateCircleGraphicSVG(WidthRatio, ""))),
				app.Div().Class("mdc-circular-progress__gap-patch").Body(app.Raw(c.IndeterminateCircleGraphicSVG(WidthRatioSmall, ""))),
				app.Div().Class("mdc-circular-progress__circle-clipper", "mdc-circular-progress__circle-right").Body(app.Raw(c.IndeterminateCircleGraphicSVG(WidthRatio, ""))),
			),
		)
	} else {
		var fourColorsBody []app.UI
		for i, color := range c.colors {
			layer := app.Div().Class("mdc-circular-progress__spinner-layer", fmt.Sprintf("mdc-circular-progress__color-%d", i+1)).Body(
				app.Div().Class("mdc-circular-progress__circle-clipper", "mdc-circular-progress__circle-left").Body(app.Raw(c.IndeterminateCircleGraphicSVG(WidthRatio, color))),
				app.Div().Class("mdc-circular-progress__gap-patch").Body(app.Raw(c.IndeterminateCircleGraphicSVG(WidthRatioSmall, color))),
				app.Div().Class("mdc-circular-progress__circle-clipper", "mdc-circular-progress__circle-right").Body(app.Raw(c.IndeterminateCircleGraphicSVG(WidthRatio, color))),
			)
			fourColorsBody = append(fourColorsBody, layer)
		}
		indeterminateContainer.Body(fourColorsBody...)
	}

	root.Body(determinateContainer, indeterminateContainer)

	return root
}

const RadiusRatio = float64(18) / float64(48)
const WidthRatio = float64(4) / float64(48)
const WidthRatioSmall = float64(3.2) / float64(48)

func (c *Circular) DeterminateSVG() string {

	size := float64(c.size)
	center := c.size / 2
	radius := size * RadiusRatio
	//width := size * WidthRatio
	// circumference is 2*PI*r
	stroke := 2 * math.Pi * radius
	strokeWidth := size * WidthRatio

	result := fmt.Sprintf(`<svg class="mdc-circular-progress__determinate-circle-graphic" viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg">\n`, c.size, c.size)
	result += fmt.Sprintf(`<circle class="mdc-circular-progress__determinate-track" cx="%d" cy="%d" r="%.1f" stroke-width="%.1f"/>\n`, center, center, radius, strokeWidth)
	result += fmt.Sprintf(`<circle class="mdc-circular-progress__determinate-circle" cx="%d" cy="%d" r="%.1f" stroke-dasharray="%.1f" stroke-dashoffset="%.1f" stroke-width="%.1f"/>\n`,
		center, center, radius, stroke, stroke, strokeWidth)
	result += "</svg>"
	return result
}

func (c *Circular) IndeterminateCircleGraphicSVG(ratio float64, color string) string {
	size := float64(c.size)
	center := c.size / 2
	radius := size * RadiusRatio
	// circumference is 2*PI*r
	stroke := 2 * math.Pi * radius
	strokeWidth := size * ratio

	result := fmt.Sprintf(`<svg class="mdc-circular-progress__indeterminate-circle-graphic" viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg" style="stroke: %s;">`, c.size, c.size, color)
	result += fmt.Sprintf(`<circle cx="%d" cy="%d" r="%.1f" stroke-dasharray="%.1f" stroke-dashoffset="%.1f" stroke-width="%.1f"/>`, center, center, radius, stroke, stroke/2, strokeWidth)
	result += "</svg>"
	return result
}

func (c *Circular) OnMount(ctx app.Context) {
	rootElement := app.Window().GetElementByID(c.id)
	c.target = c.JsNewAtPath("mdc.circularProgress.MDCCircularProgress", rootElement)

	if c.onmount != nil {
		c.onmount(ctx)
	}
}

func (c *Circular) Open() {
	if !c.target.IsUndefined() {
		c.target.Call("open")
	}
}

func (c *Circular) Determinate(d bool) {
	if !c.target.IsUndefined() {
		c.target.Set("determinate", d)
	}
}

func (c *Circular) SetProgress(f float64) {
	c.progress = f
	if !c.target.IsUndefined() {
		c.target.Set("progress", c.progress)
	}
}
func (c *Circular) Close() {
	if !c.target.IsUndefined() {
		c.target.Call("close")
	}
}
