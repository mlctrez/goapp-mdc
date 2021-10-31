package list

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
)

func TestItem_Render(t *testing.T) {
	compo := &Item{Text: "primary"}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Li().Class("mdc-deprecated-list-item"))
	ct.At(0, 0).Match(app.Span().Class("mdc-deprecated-list-item__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-deprecated-list-item__text"))
	ct.At(0, 1, 0).Match(app.Text(compo.Text))
}

func TestItem_Render_Secondary(t *testing.T) {
	compo := &Item{Text: "primary", Secondary: "secondary"}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Li().Class("mdc-deprecated-list-item"))
	ct.At(0, 0).Match(app.Span().Class("mdc-deprecated-list-item__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-deprecated-list-item__text"))
	ct.At(0, 1, 0).Match(app.Span().Class("mdc-deprecated-list-item__primary-text"))
	ct.At(0, 1, 0, 0).Match(app.Text(compo.Text))
	ct.At(0, 1, 1).Match(app.Span().Class("mdc-deprecated-list-item__secondary-text"))
	ct.At(0, 1, 1, 0).Match(app.Text(compo.Secondary))
}

func TestItem_Render_Option(t *testing.T) {
	compo := &Item{Text: "one", Type: ItemTypeOption}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Li().Class("mdc-deprecated-list-item").Attr("role", "option"))
	ct.At(0, 0).Match(app.Span().Class("mdc-deprecated-list-item__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-deprecated-list-item__text"))
	ct.At(0, 1, 0).Match(app.Text(compo.Text))
}

//func TestItem_Render_Radio(t *testing.T) {
//	compo := &Item{Text: "one", Type: ItemTypeRadio}
//	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
//	defer ct.Close()
//
//	ct.At(0).Match(app.Li().Class("mdc-deprecated-list-item").Attr("role", "radio"))
//	ct.At(0, 0).Match(app.Span().Class("mdc-deprecated-list-item__ripple"))
//	ct.At(0, 1).Match(app.Span().Class("mdc-deprecated-list-item__text"))
//	ct.At(0, 1, 0).Match(app.Text(compo.Text))
//}

func TestItem_Render_Checkbox(t *testing.T) {
	compo := &Item{Text: "one", Type: ItemTypeCheckbox, id: "testing"}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Li().Class("mdc-deprecated-list-item").Attr("role", "checkbox"))
	ct.At(0, 0).Match(app.Span().Class("mdc-deprecated-list-item__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-deprecated-list-item__graphic"))
	ct.At(0, 1,0).Match(app.Div().Class("mdc-checkbox"))
	ct.At(0, 1,0,0).Match(app.Input().Type("checkbox").Class("mdc-checkbox__native-control").ID(compo.id))
	ct.At(0, 1,0,1).Match(checkbox.MDCCheckboxBackground())
	ct.At(0, 2).Match(app.Label().Class("mdc-deprecated-list-item__text").For(compo.id).Text(compo.Text))

}

func TestItem_Render_ItemSelectStateTabZero(t *testing.T) {
	compo := &Item{Text: "one", state: ItemSelectStateTabZero}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Li().Class("mdc-deprecated-list-item").TabIndex(0))
	ct.At(0, 0).Match(app.Span().Class("mdc-deprecated-list-item__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-deprecated-list-item__text"))
	ct.At(0, 1, 0).Match(app.Text(compo.Text))
}

func TestItem_Render_ItemSelectStateSelected(t *testing.T) {
	compo := &Item{Text: "one", state: ItemSelectStateSelected}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	// removing FOUC one test at a time
	ct.At(0).Match(app.Li().Class("mdc-deprecated-list-item mdc-deprecated-list-item--selected").
		TabIndex(0).Aria("selected", "true"))
	ct.At(0, 0).Match(app.Span().Class("mdc-deprecated-list-item__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-deprecated-list-item__text"))
	ct.At(0, 1, 0).Match(app.Text(compo.Text))
}

func TestItem_Render_ItemSelectStateNotSelected(t *testing.T) {
	compo := &Item{Text: "one", state: ItemSelectStateNotSelected}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	// removing FOUC one test at a time
	ct.At(0).Match(app.Li().Class("mdc-deprecated-list-item").Aria("selected", "false"))
	ct.At(0, 0).Match(app.Span().Class("mdc-deprecated-list-item__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-deprecated-list-item__text"))
	ct.At(0, 1, 0).Match(app.Text(compo.Text))
}
