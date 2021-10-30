package icon

import (
	"testing"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

func TestButton_Render_Plain(t *testing.T) {

	compo := &Button{Icon: MIBookmark, AriaOff: "press this button"}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Button().Class("mdc-icon-button material-icons").Aria("label", compo.AriaOff))
	ct.At(0, 0).Match(app.Div().Class("mdc-icon-button__ripple"))
	ct.At(0, 1).Match(app.Text(compo.Icon))

}

func TestButton_Render_Toggle_Off(t *testing.T) {

	compo := &Button{
		Id:      uuid.New().String(),
		Icon:    MIFavorite,
		IconOff: MIFavoriteBorder,
		AriaOff: "add to favorites",
		AriaOn:  "remove from favorites",
	}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(
		app.Button().ID(compo.Id).Class("mdc-icon-button").Aria("label", compo.AriaOff).
			DataSet("aria-label-on", compo.AriaOn).DataSet("aria-label-off", compo.AriaOff))
	ct.At(0, 0).Match(app.Div().Class("mdc-icon-button__ripple"))
	ct.At(0, 1).Match(compo.Icon.I().Class("mdc-icon-button__icon mdc-icon-button__icon--on"))
	ct.At(0, 2).Match(compo.IconOff.I().Class("mdc-icon-button__icon"))

}

func TestButton_Render_Toggle_On(t *testing.T) {

	compo := &Button{
		Id:      uuid.New().String(),
		Icon:    MIFavorite,
		IconOff: MIFavoriteBorder,
		State:   true,
		AriaOn:  "remove from favorites",
		AriaOff: "add to favorites",
	}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(
		app.Button().ID(compo.Id).Class("mdc-icon-button mdc-icon-button--on").Aria("label", compo.AriaOn).
			DataSet("aria-label-on", compo.AriaOn).DataSet("aria-label-off", compo.AriaOff))
	ct.At(0, 0).Match(app.Div().Class("mdc-icon-button__ripple"))
	ct.At(0, 1).Match(compo.Icon.I().Class("mdc-icon-button__icon mdc-icon-button__icon--on"))
	ct.At(0, 2).Match(compo.IconOff.I().Class("mdc-icon-button__icon"))

}
