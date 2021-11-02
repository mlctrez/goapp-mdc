package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type CheckboxDemo struct {
	app.Compo
	checkboxOne *checkbox.Checkbox
}

func (d *CheckboxDemo) onChange(changed func(bool)) func(checkbox app.HTMLInput) {
	return func(checkbox app.HTMLInput) {
		checkbox.OnChange(func(ctx app.Context, e app.Event) {
			changed(ctx.JSSrc().Get("checked").Bool())
			d.checkboxOne.Update()
		})
	}
}

func (d *CheckboxDemo) Render() app.UI {

	if d.checkboxOne == nil {
		d.checkboxOne = &checkbox.Checkbox{Id: "checkboxOne", Label: "A Checkbox",
			Callback: func(input app.HTMLInput) {
				input.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Checked = ctx.JSSrc().Get("checked").Bool()
					if d.checkboxOne.Indeterminate {
						d.checkboxOne.Indeterminate = false
					}
					d.Update()
				})
			}}
	}

	body := layout.Grid().Body(layout.Inner().Body(
		layout.CellModified("middle", 12).
			Text("Demonstration of interacting with checkbox state from other checkboxes."),
		layout.Cell().Body(d.checkboxOne),
		layout.Cell().Body(
			&checkbox.Checkbox{Id: "checked", Label: "checked", Checked: d.checkboxOne.Checked,
				Callback: d.onChange(func(b bool) { d.checkboxOne.Checked = b })},
			&checkbox.Checkbox{Id: "indeterminate", Label: "indeterminate", Checked: d.checkboxOne.Indeterminate,
				Callback: d.onChange(func(b bool) { d.checkboxOne.Indeterminate = b })},
			&checkbox.Checkbox{Id: "disabled", Label: "disabled", Checked: d.checkboxOne.Disabled,
				Callback: d.onChange(func(b bool) { d.checkboxOne.Disabled = b })},
		),
	))

	return PageBody(body)

}
