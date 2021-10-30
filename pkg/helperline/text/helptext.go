package text

import (


	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type HelperText struct {
	app.Compo
	base.JsUtil
	id   string
	text string
}

func New(id, text string) *HelperText {
	return &HelperText{id: id, text: text}
}

func (t *HelperText) Render() app.UI {
	return app.Div().Class("mdc-text-field-helper-text").Aria("hidden", "true").Text(t.text)
}

func (t *HelperText) OnMount(ctx app.Context) {
	t.JsNewAtPath("mdc.textField.MDCTextFieldHelperText", app.Window().GetElementByID(t.id))
}
