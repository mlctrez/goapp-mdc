package card

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type Demo struct {
	app.Compo
}

func (d Demo) Render() app.UI {

	buttonCallback := func(text string) func(button app.HTMLButton) {
		return func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				fmt.Println("you clicked on button" + text)
			})
		}
	}

	return layout.Grid().Body(
		layout.Inner().Body(
			layout.Cell().Body(
				&Card{Id: uuid.New().String(), Padding: 16,
					PrimaryAction: []app.UI{app.Div().Text("Primary action card no outline")}},
			),
			layout.Cell().Body(
				&Card{Id: uuid.New().String(), Width: 200, Height: 200, Outlined: true,
					PrimaryAction: []app.UI{app.Div().Text("Primary action card 200x200px with outline")}},
			),
			layout.Cell().Body(
				&Card{Id: uuid.New().String(), Outlined: true, Padding: 16,
					PrimaryAction: []app.UI{app.Div().Text("Primary action card card with buttons")},
					ActionButtons: []app.UI{
						&button.Button{Id: uuid.New().String(), CardAction: true, Label: "Button One", Callback: buttonCallback("one")},
						&button.Button{Id: uuid.New().String(), CardAction: true, Label: "Button Two", Callback: buttonCallback("two")},
					},
				},
			),
			layout.Cell().Body(
				// an example media card with title
				&Card{Id: uuid.New().String(), Height: 100, Width: 100,
					PrimaryAction: []app.UI{
						&Media{Width: 100, Height: 100, Image: "/web/logo-192.png", Title: "Media"},
					}},
			),
			layout.Cell().Body(
				// an example media card with title
				&Card{Id: uuid.New().String(), Height: 100, Width: 100,
					PrimaryAction: []app.UI{
						&Media{Width: 100, Height: 100, Image: "/web/logo-192.png"},
					}},
			),
		),
	)
}
