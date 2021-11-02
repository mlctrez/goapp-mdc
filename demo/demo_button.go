package demo

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type ButtonDemo struct {
	app.Compo
	button *button.Button
}

func (d *ButtonDemo) Render() app.UI {
	if d.button == nil {
		d.button = &button.Button{Id: "subjectButton", Label: "a button"}
	}
	handleCheckboxChange := func(before func(checkVal bool)) func(input app.HTMLInput) {
		return func(input app.HTMLInput) {
			input.OnChange(func(ctx app.Context, e app.Event) {
				before(ctx.JSSrc().Get("checked").Bool())
				d.button.Update()
				// attempt to re-attach ripple, trying with a delay
				ctx.After(500*time.Millisecond, func(context app.Context) {
					d.button.OnMount(context)
				})
			})
		}
	}

	body := layout.Grid().Body(layout.Inner().Body(
		layout.CellModified("middle", 12).Body(d.button),
		layout.Cell().Body(
			&checkbox.Checkbox{Id: "toggleIcon", Label: "has icon",
				Callback: handleCheckboxChange(func(checkVal bool) {
					if checkVal {
						d.button.Icon = "bookmark"
					} else {
						d.button.Icon = ""
					}
				})},
			&checkbox.Checkbox{Id: "toggleTrailing", Label: "trailing icon",
				Callback: handleCheckboxChange(func(checkVal bool) { d.button.TrailingIcon = checkVal })},
			&checkbox.Checkbox{Id: "toggleOutline", Label: "outlined",
				Callback: handleCheckboxChange(func(checkVal bool) { d.button.Outlined = checkVal })},
			&checkbox.Checkbox{Id: "toggleRaised", Label: "raised",
				Callback: handleCheckboxChange(func(checkVal bool) { d.button.Raised = checkVal })},
			&checkbox.Checkbox{Id: "toggleUnelevated", Label: "unelevated",
				Callback: handleCheckboxChange(func(checkVal bool) { d.button.Unelevated = checkVal })}),
	))
	return PageBody(body)
}
