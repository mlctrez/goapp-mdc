package list

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

func TestList_Render(t *testing.T) {
	compo := &List{}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Ul().Class("mdc-deprecated-list"))
}

func TestList_Render_TwoLine(t *testing.T) {
	compo := &List{TwoLine: true}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Ul().Class("mdc-deprecated-list mdc-deprecated-list--two-line"))
}

func TestList_Render_SingleSelection(t *testing.T) {
	id := "testID"
	compo := &List{Id: id, Type: SingleSelection}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Ul().ID(compo.Id).Class("mdc-deprecated-list").Attr("role", "listbox"))
}

func TestList_Render_Navigation(t *testing.T) {
	compo := &List{Type: Navigation}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Nav().Class("mdc-deprecated-list"))
}
