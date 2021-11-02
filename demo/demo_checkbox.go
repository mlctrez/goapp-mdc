package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
)

type CheckboxDemo struct {
	app.Compo
	checkboxOne *checkbox.Checkbox
}

func (d *CheckboxDemo) Render() app.UI {

	if d.checkboxOne == nil {
		d.checkboxOne = &checkbox.Checkbox{Id: "checkboxOne", Label: "Checkbox Label",
			Callback: func(input app.HTMLInput) {
				input.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Checked = ctx.JSSrc().Get("checked").Bool()
					// changing the state of the checkbox should clear this flag
					if d.checkboxOne.Indeterminate {
						d.checkboxOne.Indeterminate = false
					}
					d.Update()
				})
			}}
	}

	body := app.Div().Body(
		d.checkboxOne,
		app.Hr(),
		&checkbox.Checkbox{Id: "checked", Label: "checked", Checked: d.checkboxOne.Checked,
			Callback: func(checkbox app.HTMLInput) {
				checkbox.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Checked = ctx.JSSrc().Get("checked").Bool()
					d.checkboxOne.Update()
				})
			}},
		&checkbox.Checkbox{Id: "indeterminate", Label: "indeterminate", Checked: d.checkboxOne.Indeterminate,
			Callback: func(checkbox app.HTMLInput) {
				checkbox.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Indeterminate = ctx.JSSrc().Get("checked").Bool()
					d.checkboxOne.Update()
				})
			}},
		&checkbox.Checkbox{Id: "disabled", Label: "disabled", Checked: d.checkboxOne.Disabled,
			Callback: func(checkbox app.HTMLInput) {
				checkbox.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Disabled = ctx.JSSrc().Get("checked").Bool()
					d.checkboxOne.Update()
				})
			}},
	)
	return PageBody(body)

}
