package demo

import (
	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/helperline"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/textarea"
	"github.com/mlctrez/goapp-mdc/pkg/textfield"
)

type FormDemo struct {
	app.Compo
}

func id() string {
	return uuid.New().String()
}

func textAreaExample() []app.UI {
	idOne := id()
	taOne := textarea.New(idOne).Size(8, 40).Outlined(true).
		Label("outlined text area").MaxLength(240)
	helpOne := helperline.New(idOne, "textarea help text", "0 / 240")

	return []app.UI{app.Div().Style("display", "inline-block").Body(taOne, helpOne)}
}

func (e *FormDemo) Render() app.UI {

	//body := app.Div().Style("display", "block").Body(
	//	app.Div().Text("test1"),
	//	app.Div().Text("test2"),
	//)

	body := layout.Grid().Body(layout.Inner().Style("display", "flex").Body(
		layout.Cell().Body(layout.Inner().Style("display", "flex").Body(
			layout.CellWide().Body(app.H4().Text("Text Area")),
			layout.Cell().Body(&textfield.TextField{Id: id(), Label: "normal"}),
			layout.Cell().Body(&textfield.TextField{Id: id(), Label: "required", Required: true}),
			layout.Cell().Body(&textfield.TextField{Id: id(), Label: "outlined", Outlined: true}),
			layout.Cell().Body(&textfield.TextField{Id: id(), Label: "outlined required",
				Outlined: true, Required: true}),
			layout.Cell().Body(&textfield.TextField{Id: id(), Placeholder: "placeholder"}),
		)),
		layout.Cell().Body(layout.Inner().Style("display", "flex").Body(
			layout.CellWide().Body(app.H4().Text("Text Field")),
			layout.Cell().Body(textAreaExample()...),
		)),
	))

	_ = body
	return PageBody(body)
}
