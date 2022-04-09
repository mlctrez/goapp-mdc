package dialog

import (
	"testing"

	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/button"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func TestDialog_Render(t *testing.T) {
	id := "dialogId"
	compo := &Dialog{Id: id}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().Class("mdc-dialog").DataSet(autoinit.MdcAutoInitData, API).ID(compo.Id))
	ct.At(0, 0).Match(app.Div().Class("mdc-dialog__container"))
	ct.At(0, 0, 0).Match(app.Div().Class("mdc-dialog__surface").
		Attr("role", "alertdialog").Aria("modal", "true"))
	ct.At(0, 0, 0, 0).Match(app.Div().Class("mdc-dialog__actions"))
	ct.At(0, 1).Match(app.Div().Class("mdc-dialog__scrim"))

}

func TestDialog_Render_Button(t *testing.T) {
	id := "dialogId"
	compo := &Dialog{Id: id}
	okButton := &button.Button{Id: "okButton"}
	cancelButton := &button.Button{Id: "cancelButton"}
	compo.Buttons = []app.UI{okButton, cancelButton}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().Class("mdc-dialog").DataSet(autoinit.MdcAutoInitData, API).ID(compo.Id))
	ct.At(0, 0).Match(app.Div().Class("mdc-dialog__container"))
	ct.At(0, 0, 0).Match(app.Div().Class("mdc-dialog__surface").
		Attr("role", "alertdialog").Aria("modal", "true"))
	ct.At(0, 0, 0, 0).Match(app.Div().Class("mdc-dialog__actions"))
	ct.At(0, 0, 0, 0, 0).Match(okButton)
	ct.At(0, 0, 0, 0, 1).Match(cancelButton)
	ct.At(0, 1).Match(app.Div().Class("mdc-dialog__scrim"))

}

func TestDialog_Render_Button_And_Content(t *testing.T) {
	id := "dialogId"
	compo := &Dialog{Id: id}
	okButton := &button.Button{Id: "okButton"}
	cancelButton := &button.Button{Id: "cancelButton"}
	compo.Buttons = []app.UI{okButton, cancelButton}
	content := app.Span().Text("Some content here")
	compo.Content = []app.UI{content}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().Class("mdc-dialog").DataSet(autoinit.MdcAutoInitData, API).ID(compo.Id))
	ct.At(0, 0).Match(app.Div().Class("mdc-dialog__container"))
	ct.At(0, 0, 0).Match(
		app.Div().Class("mdc-dialog__surface").Attr("role", "alertdialog").
			Aria("modal", "true").Aria("describedby", compo.Id+"-content"),
	)
	ct.At(0, 0, 0, 0).Match(app.Div().Class("mdc-dialog__content").ID(compo.Id + "-content"))
	ct.At(0, 0, 0, 0, 0).Match(content)
	ct.At(0, 0, 0, 1).Match(app.Div().Class("mdc-dialog__actions"))
	ct.At(0, 0, 0, 1, 0).Match(okButton)
	ct.At(0, 0, 0, 1, 1).Match(cancelButton)
	ct.At(0, 1).Match(app.Div().Class("mdc-dialog__scrim"))

}

func TestDialog_Render_All(t *testing.T) {
	id := "dialogId"
	compo := &Dialog{Id: id}
	okButton := &button.Button{Id: "okButton"}
	cancelButton := &button.Button{Id: "cancelButton"}
	compo.Buttons = []app.UI{okButton, cancelButton}
	content := app.Span().Text("Some content here")
	compo.Content = []app.UI{content}
	title := app.Span().Text("Title")
	compo.Title = []app.UI{title}

	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().Class("mdc-dialog").DataSet(autoinit.MdcAutoInitData, API).ID(compo.Id))
	ct.At(0, 0).Match(app.Div().Class("mdc-dialog__container"))
	ct.At(0, 0, 0).Match(
		app.Div().Class("mdc-dialog__surface").Attr("role", "alertdialog").
			Aria("modal", "true").Aria("describedby", compo.Id+"-content").
			Aria("labelledby", compo.Id+"-title"),
	)
	ct.At(0, 0, 0, 0).Match(app.Div().Class("mdc-dialog__title").ID(compo.Id + "-title"))
	ct.At(0, 0, 0, 0, 0).Match(title)
	ct.At(0, 0, 0, 1).Match(app.Div().Class("mdc-dialog__content").ID(compo.Id + "-content"))
	ct.At(0, 0, 0, 1, 0).Match(content)
	ct.At(0, 0, 0, 2).Match(app.Div().Class("mdc-dialog__actions"))
	ct.At(0, 0, 0, 2, 0).Match(okButton)
	ct.At(0, 0, 0, 2, 1).Match(cancelButton)
	ct.At(0, 1).Match(app.Div().Class("mdc-dialog__scrim"))

}
