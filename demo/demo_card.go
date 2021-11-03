package demo

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/card"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type CardDemo struct {
	app.Compo
}

func (d CardDemo) Render() app.UI {

	buttonCallback := func(text string) func(button app.HTMLButton) {
		return func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				fmt.Println("you clicked on button" + text)
			})
		}
	}

	body := FlexGrid(
		layout.Cell().Body(
			&card.Card{Id: uuid.New().String(), Padding: 16,
				PrimaryAction: []app.UI{app.Div().Text("Primary action card no outline")}},
		),
		layout.Cell().Body(
			&card.Card{Id: uuid.New().String(), Width: 200, Height: 200, Outlined: true,
				PrimaryAction: []app.UI{app.Div().Text("Primary action card 200x200px with outline")}},
		),
		layout.Cell().Body(
			&card.Card{Id: uuid.New().String(), Outlined: true, Padding: 16,
				PrimaryAction: []app.UI{app.Div().Text("Primary action card card with buttons")},
				ActionButtons: []app.UI{
					&button.Button{Id: uuid.New().String(), CardAction: true,
						Label: "Button One", Callback: buttonCallback("one")},
					&button.Button{Id: uuid.New().String(), CardAction: true,
						Label: "Button Two", Callback: buttonCallback("two")},
				},
			},
		),
		layout.Cell().Body(GopherCard("Media")),
		layout.Cell().Body(GopherCard("")),
		gopherAttribution(),
	)

	return PageBody(body)
}


func gopherAttribution() app.HTMLDiv {
	return layout.CellModified("bottom", 12).Body(
		app.Text("Gopher images courtesy of "),
		app.A().Href("https://github.com/golang-samples/gopher-vector").Text("gopher-vector"),
		app.Br(),
		app.Text("Licensed under the Creative Commons 3.0 Attributions license."))
}

func GopherCard(title string) app.UI {
	return &card.Card{Id: uuid.New().String(), Width: 202, Height: 259,
		PrimaryAction: []app.UI{&card.Media{
			Width: 202, Height: 259, Image: "/web/gopher-front.png", Title: title}}}
}
