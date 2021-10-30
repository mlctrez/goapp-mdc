package banner

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

func TestBanner_Render(t *testing.T) {
	id := "testID"
	compo := &Banner{Id: id, Text: "this is the banner text"}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().Class("mdc-banner").ID(compo.Id))
	ct.At(0, 0).Match(app.Div().Class("mdc-banner__content").
		Attr("role", "alertdialog").Aria("live", "assertlive"))
	ct.At(0, 0, 0).Match(app.Div().Class("mdc-banner__graphic-text-wrapper"))
	ct.At(0, 0, 0, 0).Match(app.Div().Class("mdc-banner__text").Text(compo.Text + "foo"))

}

func TestBanner_Center(t *testing.T) {
	id := "testID"
	compo := &Banner{Id: id, Text: "this is the banner text", Centered: true}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().Class("mdc-banner mdc-banner--centered").ID(compo.Id))
	ct.At(0, 0).Match(app.Div().Class("mdc-banner__content").
		Attr("role", "alertdialog").Aria("live", "assertlive"))
	ct.At(0, 0, 0).Match(app.Div().Class("mdc-banner__graphic-text-wrapper"))
	ct.At(0, 0, 0, 0).Match(app.Div().Class("mdc-banner__text").Text(compo.Text + "foo"))

}

func TestBanner_Fixed(t *testing.T) {
	id := "testID"
	compo := &Banner{Id: id, Text: "this is the banner text", Fixed: true}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().Class("mdc-banner").ID(compo.Id))
	ct.At(0, 0).Match(app.Div().Class("mdc-banner__fixed"))
	ct.At(0, 0, 0).Match(app.Div().Class("mdc-banner__content").
		Attr("role", "alertdialog").Aria("live", "assertlive"))
	ct.At(0, 0, 0, 0).Match(app.Div().Class("mdc-banner__graphic-text-wrapper"))
	ct.At(0, 0, 0, 0, 0).Match(app.Div().Class("mdc-banner__text").Text(compo.Text + "foo"))

}
