package textfield

import (
	"testing"

	"github.com/mlctrez/goapp-mdc/pkg/base"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func TestTextField_Render(t *testing.T) {

	id := "textboxId"
	compo := &TextField{Id: id}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Label().ID(id).Class("mdc-text-field mdc-text-field--filled mdc-text-field--no-label"))
	ct.At(0, 0).Match(app.Span().Class("mdc-text-field__ripple"))
	ct.At(0, 1).Match(app.Input().ID(id + "-input").Class("mdc-text-field__input").Type("text").Value(""))
	ct.At(0, 2).Match(app.Span().Class("mdc-line-ripple"))

}

func TestTextField_Render_Outline(t *testing.T) {

	id := "textboxId"
	compo := &TextField{Id: id, Outlined: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Label().ID(id).Class("mdc-text-field mdc-text-field--outlined mdc-text-field--no-label"))
	ct.At(0, 0).Match(app.Span().Class("mdc-notched-outline"))
	ct.At(0, 0, 0).Match(app.Span().Class("mdc-notched-outline__leading"))
	ct.At(0, 0, 1).Match(app.Span().Class("mdc-notched-outline__trailing"))
	ct.At(0, 1).Match(app.Input().ID(id + "-input").Class("mdc-text-field__input").Type("text").Value(""))

}

func TestTextField_Render_Outline_And_Label(t *testing.T) {

	id := "textboxId"
	compo := &TextField{Id: id, Outlined: true, Label: "the label"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Label().ID(id).Class("mdc-text-field mdc-text-field--outlined"))
	ct.At(0, 0).Match(app.Span().Class("mdc-notched-outline"))
	ct.At(0, 0, 0).Match(app.Span().Class("mdc-notched-outline__leading"))
	ct.At(0, 0, 1).Match(app.Span().Class("mdc-notched-outline__notch"))
	ct.At(0, 0, 1, 0).Match(app.Span().Class("mdc-floating-label").
		ID(id + "-label").Text(compo.Label))
	ct.At(0, 0, 2).Match(app.Span().Class("mdc-notched-outline__trailing"))
	ct.At(0, 1).Match(app.Input().ID(id+"-input").Class("mdc-text-field__input").
		Type("text").Value("").Aria("labelledby", id+"-label"))

}

func TestTextField_Render_Value(t *testing.T) {

	id := "textboxId"
	compo := &TextField{Id: id, Value: "foo"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Label().ID(id).Class("mdc-text-field mdc-text-field--filled mdc-text-field--no-label"))
	ct.At(0, 0).Match(app.Span().Class("mdc-text-field__ripple"))
	ct.At(0, 1).Match(app.Input().ID(id + "-input").Class("mdc-text-field__input").Type("text").Value("foo"))
	ct.At(0, 2).Match(app.Span().Class("mdc-line-ripple"))

}

func TestTextField_Render_Label(t *testing.T) {

	id := "textboxId"
	compo := &TextField{Id: id, Label: "the label"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Label().ID(id).Class("mdc-text-field mdc-text-field--filled"))
	ct.At(0, 0).Match(app.Span().Class("mdc-text-field__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-floating-label").ID(id + "-label").Text(compo.Label))
	ct.At(0, 2).Match(
		app.Input().ID(id+"-input").Class("mdc-text-field__input").Aria("labelledby", id+"-label").
			Type("text").Value(compo.Value))
	ct.At(0, 3).Match(app.Span().Class("mdc-line-ripple"))

}

func TestTextField_Render_Label_And_Value(t *testing.T) {

	id := "textboxId"
	compo := &TextField{Id: id, Label: "the label", Value: "the value"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Label().ID(id).Class("mdc-text-field mdc-text-field--filled mdc-text-field--label-floating"))
	ct.At(0, 0).Match(app.Span().Class("mdc-text-field__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-floating-label mdc-floating-label mdc-floating-label--float-above").ID(id + "-label").Text(compo.Label))
	ct.At(0, 2).Match(
		app.Input().ID(id+"-input").Class("mdc-text-field__input").Aria("labelledby", id+"-label").
			Type("text").Value(compo.Value))
	ct.At(0, 3).Match(app.Span().Class("mdc-line-ripple"))

}

func TestTextField_Render_Required(t *testing.T) {

	id := "textboxId"
	compo := &TextField{Id: id, Required: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Label().ID(id).Class("mdc-text-field mdc-text-field--filled mdc-text-field--no-label"))
	ct.At(0, 0).Match(app.Span().Class("mdc-text-field__ripple"))
	ct.At(0, 1).Match(
		app.Input().ID(id + "-input").Class("mdc-text-field__input").
			Type("text").Value("").Required(true))
	ct.At(0, 2).Match(app.Span().Class("mdc-line-ripple"))

}

func TestTextField_Render_Placeholder(t *testing.T) {

	id := "textboxId"
	compo := &TextField{Id: id, Placeholder: "placeholder"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Label().ID(id).Class("mdc-text-field mdc-text-field--filled mdc-text-field--no-label"))
	ct.At(0, 0).Match(app.Span().Class("mdc-text-field__ripple"))
	ct.At(0, 1).Match(
		app.Input().ID(id+"-input").Class("mdc-text-field__input").Type("text").Value(compo.Value).
			Placeholder(compo.Placeholder).Aria("label", "Label"))
	ct.At(0, 2).Match(app.Span().Class("mdc-line-ripple"))

}
