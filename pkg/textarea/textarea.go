package textarea

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type TextArea struct {
	app.Compo
	base.JsUtil
	id          string
	label       string
	placeholder string
	outlined    bool
	value       string
	resizable   bool
	rows        int
	cols        int
	maxlength   int
	callback    func(in app.HTMLTextarea)
	onmount     func(ctx app.Context)
}

func New(id string) *TextArea {
	return &TextArea{id: id, rows: 8, cols: 40}
}

func (t *TextArea) Label(label string) *TextArea {
	t.label = label
	return t
}

func (t *TextArea) Placeholder(placeholder string) *TextArea {
	t.placeholder = placeholder
	return t
}

func (t *TextArea) Outlined(outlined bool) *TextArea {
	t.outlined = outlined
	return t
}

func (t *TextArea) Value(val string) *TextArea {
	t.value = val
	return t
}

func (t *TextArea) Resizeable(r bool) *TextArea {
	t.resizable = r
	return t
}
func (t *TextArea) Size(rows, cols int) *TextArea {
	t.rows = rows
	t.cols = cols
	return t
}

func (t *TextArea) MaxLength(l int) *TextArea {
	t.maxlength = l
	return t
}

func (t *TextArea) WithCallback(cb func(in app.HTMLTextarea)) *TextArea {
	t.callback = cb
	return t
}

func (t *TextArea) WithOnMount(cb func(ctx app.Context)) *TextArea {
	t.onmount = cb
	return t
}

func (t *TextArea) buildClasses() []string {
	classes := []string{"mdc-text-field"}
	if t.outlined {
		classes = append(classes, "mdc-text-field--outlined")
	} else {
		classes = append(classes, "mdc-text-field--filled")
	}

	classes = append(classes, "mdc-text-field--textarea")

	if t.label == "" {
		classes = append(classes, "mdc-text-field--no-label")
	} else {
		if t.value != "" {
			classes = append(classes, "mdc-text-field--label-floating")
		}
	}
	return classes
}

func (t *TextArea) Render() app.UI {

	classes := t.buildClasses()
	var input app.HTMLTextarea
	var labelSpan app.HTMLSpan

	input = app.Textarea().Class("mdc-text-field__input").ID(t.id + "-input").Text(t.value)
	input.MaxLength(t.maxlength).Rows(t.rows).Cols(t.cols)
	if t.label == "" {
		if t.placeholder != "" {
			input.Placeholder(t.placeholder).Aria("label", "Label").Text(t.value)
		}
	} else {
		labelId := t.id + "-label"
		labelSpan = app.Span().Class("mdc-floating-label").ID(labelId).Text(t.label)
		if t.value != "" {
			labelSpan.Class("mdc-floating-label", "mdc-floating-label--float-above")
		}
		input.Aria("labelledby", labelId).ID(t.id + "-input")
	}

	if t.callback != nil {
		t.callback(input)
	}

	return app.Label().ID(t.id).Class(classes...).Body(
		app.If(t.outlined,
			app.Span().Class("mdc-notched-outline").Body(
				app.Span().Class("mdc-notched-outline__leading"),
				app.If(labelSpan != nil,
					app.Span().Class("mdc-notched-outline__notch").Body(labelSpan)),
				app.Span().Class("mdc-notched-outline__trailing"),
			),
			app.If(t.resizable, app.Span().Class("mdc-text-field__resizer").Body(input)).Else(input),
		).Else(
			app.Span().Class("mdc-text-field__ripple"),
			app.If(labelSpan != nil, labelSpan),
			app.If(t.resizable, app.Span().Class("mdc-text-field__resizer").Body(input)).Else(input),
			app.Span().Class("mdc-line-ripple"),
		),
	)
}

func (t *TextArea) OnMount(ctx app.Context) {
	value := t.JsValueAtPath("mdc.textField.MDCTextField")
	if !value.IsUndefined() {
		value.Call("attachTo", app.Window().GetElementByID(t.id))
	}
	if t.onmount != nil {
		t.onmount(ctx)
	}
}
