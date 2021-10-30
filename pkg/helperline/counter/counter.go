package counter

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type CharacterCounter struct {
	app.Compo
	base.JsUtil
	id   string
	text string
}

func New(id, text string) *CharacterCounter {
	return &CharacterCounter{id: id, text: text}
}

func (t *CharacterCounter) Render() app.UI {
	return app.Div().Class("mdc-text-field-character-counter").Text(t.text)
}

func (t *CharacterCounter) OnMount(ctx app.Context) {
	value := t.JsValueAtPath("mdc.textField.MDCTextFieldCharacterCounter")
	if !value.IsUndefined() {
		value.Call("attachTo", app.Window().GetElementByID(t.id))
	}
}
