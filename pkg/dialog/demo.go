package dialog

import (
	"fmt"
	"github.com/mlctrez/goapp-mdc/pkg/button"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Demo struct {
	app.Compo
}

func (d *Demo) Render() app.UI {
	diagId := uuid.New().String()

	diag := &Dialog{Id: diagId}
	diag.Title = []app.UI{app.Div().Text("Dialog Title")}
	diag.Content = []app.UI{app.Div().Text("This is the content section of the dialog. There is quite a bit of text here to demonstrate how the dialog renders with this amount of text.")}

	diag.Buttons = []app.UI{
		&button.Button{Id: diagId + "-cancel", Dialog: true, DialogAction: "cancel", Label: "cancel", Callback: func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				fmt.Println("you clicked on the cancel button")
			})
		}},
		&button.Button{Id: diagId + "-dismiss", Dialog: true, DialogAction: "dismiss", Label: "dismiss"},
	}

	openDialog := &button.Button{Id: "openDialogButton", Label: "open dialog"}
	openDialog.Callback = func(b app.HTMLButton) {
		b.OnClick(func(ctx app.Context, e app.Event) {
			diag.Open()
		})
	}

	return app.Div().Body(openDialog, diag)
}
