package helperline

import (
	"github.com/mlctrez/goapp-mdc/pkg/helperline/counter"
	"github.com/mlctrez/goapp-mdc/pkg/helperline/text"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type HelperLine struct {
	app.Compo
	id          string
	helpText    string
	counterText string
}

func New(id, helpText, counterText string) *HelperLine {
	return &HelperLine{id: id, helpText: helpText, counterText: counterText}
}

func (t *HelperLine) Render() app.UI {

	var elements []app.UI

	if t.helpText != "" {
		elements = append(elements, text.New(t.id+"-help", t.helpText))
	}
	if t.counterText != "" {
		elements = append(elements, counter.New(t.id+"-counter", t.helpText))
	}
	if len(elements) == 0 {
		elements = append(elements, text.New(t.id+"-help", "HelperLine used without help or counter set"))
	}

	return app.Div().Class("mdc-text-field-helper-line").Body(elements...)
}
