package textfield

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type TextField struct {
	app.Compo
	base.JsUtil
	Id          string
	Label       string
	Placeholder string
	Outlined    bool
	Value       string
	Required    bool
	OnChange    func(string)
}

func (t *TextField) buildClasses() []string {
	classes := []string{"mdc-text-field"}
	if t.Outlined {
		classes = append(classes, "mdc-text-field--outlined")
	} else {
		classes = append(classes, "mdc-text-field--filled")
	}
	if t.Label == "" {
		classes = append(classes, "mdc-text-field--no-label")
	} else {
		if t.Value != "" {
			classes = append(classes, "mdc-text-field--label-floating")
		}
	}
	return classes
}

func (t *TextField) Render() app.UI {

	classes := t.buildClasses()
	var labelSpan app.HTMLSpan

	input := app.Input().Class("mdc-text-field__input").
		Type("text").ID(t.Id + "-input").Value(t.Value).OnChange(t.ValueTo(&t.Value))

	if t.Required {
		input.Required(t.Required)
	}

	if t.Label == "" {
		if t.Placeholder != "" {
			input.Placeholder(t.Placeholder).Aria("label", "Label")
		}
	} else {
		labelId := t.Id + "-label"
		labelSpan = app.Span().Class("mdc-floating-label").ID(labelId).Text(t.Label)
		if t.Value != "" {
			labelSpan.Class("mdc-floating-label", "mdc-floating-label--float-above")
		}
		input.Aria("labelledby", labelId)
	}

	return app.Label().ID(t.Id).Class(classes...).Body(
		app.If(t.Outlined,
			app.Span().Class("mdc-notched-outline").Body(
				app.Span().Class("mdc-notched-outline__leading"),
				app.If(labelSpan != nil,
					app.Span().Class("mdc-notched-outline__notch").Body(labelSpan)),
				app.Span().Class("mdc-notched-outline__trailing"),
			),
			input,
		).Else(
			app.Span().Class("mdc-text-field__ripple"),
			app.If(labelSpan != nil, labelSpan),
			input,
			app.Span().Class("mdc-line-ripple"),
		),
	)

}

func (t *TextField) OnMount(ctx app.Context) {
	t.JsNewAtPath("mdc.textField.MDCTextField", app.Window().GetElementByID(t.Id))
}
