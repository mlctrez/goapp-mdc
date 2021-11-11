package layout

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type W struct {
	app.Compo
	ui app.UI
}

func (w *W) Render() app.UI {
	return w.ui
}

func TestGrid(t *testing.T) {
	compo := &W{ui: Grid()}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Div().Class("mdc-layout-grid"))
}

func TestInner(t *testing.T) {
	compo := &W{ui: Inner()}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Div().Class("mdc-layout-grid__inner"))
}

func TestCell(t *testing.T) {
	compo := &W{ui: Cell()}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Div().Class("mdc-layout-grid__cell"))
}

func TestCellModified(t *testing.T) {
	compo := &W{ui: CellModified("top", 12)}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Div().Class("mdc-layout-grid__cell mdc-layout-grid__cell--align-top mdc-layout-grid__cell--span-12"))

	compo = &W{ui: CellModified("middle", 4)}
	ct = base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Div().Class("mdc-layout-grid__cell mdc-layout-grid__cell--align-middle mdc-layout-grid__cell--span-4"))

	compo = &W{ui: CellModified("bottom", 2)}
	ct = base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Div().Class("mdc-layout-grid__cell mdc-layout-grid__cell--align-bottom mdc-layout-grid__cell--span-2"))

	compo = &W{ui: CellModified("invalid", -1)}
	ct = base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Div().Class("mdc-layout-grid__cell"))


}

func TestWide(t *testing.T) {
	compo := &W{ui: CellWide()}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()
	ct.At(0).Match(app.Div().Class("mdc-layout-grid__cell mdc-layout-grid__cell--align-middle mdc-layout-grid__cell--span-12"))
}