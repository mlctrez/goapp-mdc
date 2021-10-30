package dialog

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Dialog struct {
	app.Compo
	base.JsUtil
	Id      string
	Title   []app.UI
	Content []app.UI
	Buttons []app.UI
	onmount func(ctx app.Context)
	target  app.Value
}

func (d *Dialog) Render() app.UI {
	dialog := app.Div().Class("mdc-dialog").ID(d.Id)

	dialogSurface := app.Div().Class("mdc-dialog__surface").Attr("role", "alertdialog")
	dialogSurface.Aria("modal", "true")
	if d.Title != nil {
		dialogSurface.Aria("labelledby", d.Id+"-title")
	}
	if d.Content != nil {
		dialogSurface.Aria("describedby", d.Id+"-content")
	}
	dialogSurface.Body(
		app.If(d.Title != nil, app.Div().Class("mdc-dialog__title").ID(d.Id+"-title").Body(d.Title...)).Else(),
		app.If(d.Content != nil, app.Div().Class("mdc-dialog__content").ID(d.Id+"-content").Body(d.Content...)).Else(),
		app.Div().Class("mdc-dialog__actions").Body(d.Buttons...),
	)

	dialog.Body(
		app.Div().Class("mdc-dialog__container").Body(dialogSurface),
		app.Div().Class("mdc-dialog__scrim"),
	)
	return dialog
}

func (d *Dialog) OnMount(ctx app.Context) {
	d.target = d.JsNewAtPath("mdc.dialog.MDCDialog", app.Window().GetElementByID(d.Id))
	if d.onmount != nil {
		d.onmount(ctx)
	}
}

func (d *Dialog) Open() {
	if !d.target.IsUndefined() {
		d.target.Call("open")
	}
}
