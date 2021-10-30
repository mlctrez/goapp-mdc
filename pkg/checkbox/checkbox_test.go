package checkbox

import (
	"testing"

	"github.com/mlctrez/goapp-mdc/pkg/base"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func TestCheckbox_Render(t *testing.T) {
	id := "checkboxId"
	compo := &Checkbox{Id: id}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().ID(id + "-formField").Class("mdc-form-field"))
	ct.At(0, 0).Match(app.Div().ID(id).Class("mdc-checkbox"))
	ct.At(0, 0, 0).Match(
		app.Input().ID(id + "-input").Class("mdc-checkbox__native-control").
			Type("checkbox").Checked(false).Disabled(false))
	ct.At(0, 0, 1).Match(app.Div().Class("mdc-checkbox__background"))
	ct.At(0, 0, 1, 0).Match(app.Raw(SVG))
	ct.At(0, 0, 1, 1).Match(app.Div().Class("mdc-checkbox__mixedmark"))
	ct.At(0, 0, 2).Match(app.Div().Class("mdc-checkbox__ripple"))

}

func TestCheckbox_Render_Checked(t *testing.T) {
	id := "checkboxId"
	compo := &Checkbox{Id: id, Label: "", Checked: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0, 0, 0).Match(
		app.Input().ID(id + "-input").Class("mdc-checkbox__native-control").
			Type("checkbox").Checked(true).Disabled(false))

}

func TestCheckbox_Render_Label(t *testing.T) {
	id := "checkboxId"
	compo := &Checkbox{Id: id, Label: "the label"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0, 1).Match(app.Label().ID(id + "-label").For(id + "-input").Text(compo.Label))
}

func TestCheckbox_Render_Disabled(t *testing.T) {
	id := "checkboxId"
	compo := &Checkbox{Id: id, Disabled: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0, 0).Match(app.Div().ID(id).Class("mdc-checkbox mdc-checkbox--disabled"))
	ct.At(0, 0, 0).Match(
		app.Input().ID(id + "-input").Class("mdc-checkbox__native-control").
			Type("checkbox").Checked(false).Disabled(true))

}

func TestCheckbox_Render_Indeterminate(t *testing.T) {
	id := "checkboxId"
	compo := &Checkbox{Id: id, Label: "", Indeterminate: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0, 0, 0).Match(
		app.Input().ID(id+"-input").Class("mdc-checkbox__native-control").
			Type("checkbox").Checked(false).DataSet("indeterminate", "true").
			Disabled(false))

}
