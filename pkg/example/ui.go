package example

import (
	"fmt"

	"github.com/mlctrez/goapp-mdc/pkg/fab"
	"github.com/mlctrez/goapp-mdc/pkg/helperline"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/recaptcha"
	"github.com/mlctrez/goapp-mdc/pkg/textarea"
	"github.com/mlctrez/goapp-mdc/pkg/textfield"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fetch "github.com/mlctrez/wasm-fetch"
)

type Index struct {
	app.Compo
}

func (i *Index) Render() app.UI {

	return app.Div().ID("indexPage").Body(
		app.Div().Text("nothing to see here, version is "+app.Getenv("GOAPP_VERSION")).OnClick(func(ctx app.Context, e app.Event) {

			var response, err = fetch.Fetch("https://httpbin.org/json", &fetch.Opts{Method: fetch.MethodGet})
			if err != nil {
				panic(err)
			}
			if response.Status == 200 {
				fmt.Println(string(response.Body))
			}

		}),
		recaptcha.New("homeRecaptcha", "homepage"),
	)
}

type Example struct {
	app.Compo
}

func id() string {
	return uuid.New().String()
}

func cell(name string, contents ...app.UI) app.UI {
	return layout.CellModified("", 12).Body(app.H4().Text(name),
		layout.Inner().Body(
			func() []app.UI {
				var result []app.UI
				for _, content := range contents {
					result = append(result, layout.Cell().Body(content))
				}
				return result
			}()...,
		),
	)
}

func fabExamples() []app.UI {
	return []app.UI{
		&fab.Fab{Id: id(), Icon: "favorite"},
		&fab.Fab{Id: id(), Icon: "favorite", Mini: true},
		&fab.Fab{Id: id(), Icon: "favorite", Extended: true, Label: "Favorite"},
	}
}

func textFieldExamples() []app.UI {
	return []app.UI{
		&textfield.TextField{Id: id(), Label: "normal"},
		&textfield.TextField{Id: id(), Label: "required", Required: true},
		&textfield.TextField{Id: id(), Label: "outlined", Outlined: true},
		&textfield.TextField{Id: id(), Label: "outlined required", Outlined: true, Required: true},
		&textfield.TextField{Id: id(), Placeholder: "placeholder"},
	}
}

func textAreaExample() []app.UI {
	idOne := id()
	taOne := textarea.New(idOne).Size(8, 40).Outlined(true).Label("outlined text area").MaxLength(240)
	helpOne := helperline.New(idOne, "textarea help text", "0 / 240")

	return []app.UI{app.Div().Style("display", "inline-block").Body(taOne, helpOne)}

}

func (e *Example) Render() app.UI {

	return layout.Grid().Body(layout.Inner().Body(
		cell("Fab", fabExamples()...),
		cell("Text Field", textFieldExamples()...),
		cell("Text Area", textAreaExample()...),
	))

}
