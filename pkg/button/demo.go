package button

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
)

type Demo struct {
	app.Compo
	button *Button
}

func (d *Demo) Render() app.UI {

	if d.button == nil {
		d.button = &Button{Id: "subjectButton", Label: "a button"}
	}

	updateButton := func(ctx app.Context) {
		d.button.Update()
		d.button.OnMount(ctx)
	}

	return app.Div().Body(
		app.H3().Text("button demo"),
		app.Br(),
		d.button,
		app.Br(),

		&checkbox.Checkbox{Id: "toggleIcon",Label: "icon", Callback: func(input app.HTMLInput) {
			input.OnChange(func(ctx app.Context, e app.Event) {
				if ctx.JSSrc().Get("checked").Bool() {
					d.button.Icon = "bookmark"
				} else {
					d.button.Icon = ""
				}
				updateButton(ctx)
			})
		}},
		&checkbox.Checkbox{Id: "toggleTrailing",Label: "trailing icon", Callback: func(input app.HTMLInput) {
			input.OnChange(func(ctx app.Context, e app.Event) {
				d.button.TrailingIcon = ctx.JSSrc().Get("checked").Bool()
				updateButton(ctx)
			})
		}},
		&checkbox.Checkbox{Id: "toggleOutline",Label: "outlined", Callback: func(input app.HTMLInput) {
			input.OnChange(func(ctx app.Context, e app.Event) {
				d.button.Outlined = ctx.JSSrc().Get("checked").Bool()
				updateButton(ctx)
			})
		}},
		&checkbox.Checkbox{Id: "toggleRaised",Label: "raised", Callback: func(input app.HTMLInput) {
			input.OnChange(func(ctx app.Context, e app.Event) {
				d.button.Raised = ctx.JSSrc().Get("checked").Bool()
				updateButton(ctx)
			})
		}},
		&checkbox.Checkbox{Id: "toggleUnelevated",Label: "unelevated", Callback: func(input app.HTMLInput) {
			input.OnChange(func(ctx app.Context, e app.Event) {
				d.button.Unelevated = ctx.JSSrc().Get("checked").Bool()
				updateButton(ctx)
			})
		}},
	)
}
