package button

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

func TestButton_Render(t *testing.T) {

	id := "buttonId"
	compo := &Button{Id: id, Label: "normal"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().ID(compo.Id).Class("mdc-button"))
	ct.At(0, 0).Match(app.Span().Class("mdc-button__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-button__label").Text(compo.Label))

}

func TestButton_Render_Icon(t *testing.T) {

	id := "buttonId"
	compo := &Button{Id: id, Label: "normal", Icon: "bookmark"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().ID(compo.Id).Class("mdc-button mdc-button--icon-leading"))
	ct.At(0, 0).Match(app.Span().Class("mdc-button__ripple"))
	ct.At(0, 1).Match(app.I().Class("material-icons mdc-button__icon").Text(compo.Icon))
	ct.At(0, 2).Match(app.Span().Class("mdc-button__label").Text(compo.Label))

}

func TestButton_Render_IconTrail(t *testing.T) {

	id := "buttonId"
	compo := &Button{Id: id, Label: "normal", Icon: "bookmark", TrailingIcon: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().ID(compo.Id).Class("mdc-button mdc-button--icon-trailing"))
	ct.At(0, 0).Match(app.Span().Class("mdc-button__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-button__label").Text(compo.Label))
	ct.At(0, 2).Match(app.I().Class("material-icons mdc-button__icon").Text(compo.Icon))

}

func TestButton_Render_Raised(t *testing.T) {

	id := "buttonId"
	// all three are set to true here since raised will override the latter two
	compo := &Button{Id: id, Label: "normal", Raised: true, Unelevated: true, Outlined: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().ID(compo.Id).Class("mdc-button mdc-button--raised"))
	ct.At(0, 0).Match(app.Span().Class("mdc-button__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-button__label").Text(compo.Label))

}

func TestButton_Render_Unelevated(t *testing.T) {

	id := "buttonId"
	// all three are set to true here since raised will override the latter two
	compo := &Button{Id: id, Label: "normal", Raised: false, Unelevated: true, Outlined: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().ID(compo.Id).Class("mdc-button mdc-button--unelevated"))
	ct.At(0, 0).Match(app.Span().Class("mdc-button__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-button__label").Text(compo.Label))

}

func TestButton_Render_Outlined(t *testing.T) {

	id := "buttonId"
	// all three are set to true here since raised will override the latter two
	compo := &Button{Id: id, Label: "normal", Raised: false, Unelevated: false, Outlined: true}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().ID(compo.Id).Class("mdc-button mdc-button--outlined"))
	ct.At(0, 0).Match(app.Span().Class("mdc-button__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-button__label").Text(compo.Label))

}

func TestButton_Render_Dialog(t *testing.T) {

	id := "buttonId"
	compo := &Button{Id: id, Label: "normal", Dialog: true, DialogAction: "cancel"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().ID(compo.Id).Class("mdc-button mdc-dialog__button").
		DataSet("mdc-dialog-action", compo.DialogAction))
	ct.At(0, 0).Match(app.Span().Class("mdc-button__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-button__label").Text(compo.Label))

}

func TestButton_Render_Banner(t *testing.T) {

	id := "buttonId"
	compo := &Button{Id: id, Label: "primary text", Banner: true, BannerAction: "primary"}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().ID(compo.Id).Class("mdc-button mdc-banner__primary-action"))
	ct.At(0, 0).Match(app.Span().Class("mdc-button__ripple"))
	ct.At(0, 1).Match(app.Span().Class("mdc-button__label").Text(compo.Label))

}
