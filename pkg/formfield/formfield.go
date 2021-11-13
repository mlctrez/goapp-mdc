package formfield

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
)

type Component interface {
	app.UI
	LabelID() string
	MdcAPI() app.Value
}

type FormField struct {
	app.Compo
	autoinit.AutoInit
	Component Component
	Label     string
	AlignEnd  bool
	NoWrap    bool
	api       app.Value
}

const Api = autoinit.MDCFormField

func (f *FormField) Render() app.UI {
	formField := app.Div().Class("mdc-form-field")
	Api.DataMdcAutoInitDiv(formField)
	if f.AlignEnd {
		formField.Class("mdc-form-field--align-end")
	}
	if f.NoWrap {
		formField.Class("mdc-form-field--nowrap")
	}
	formField.Body(
		f.Component,
		app.Label().For(f.Component.LabelID()).Text(f.Label),
	)
	return formField
}

var _ app.Mounter = (*FormField)(nil)

func (f *FormField) OnMount(ctx app.Context) {
	f.api = f.AutoInitComponent(f.JSValue(), Api)
	f.api.Set("input", f.Component.MdcAPI())
}
