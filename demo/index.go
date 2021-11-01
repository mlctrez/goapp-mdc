package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Index struct {
	app.Compo
}

func (i *Index) Render() app.UI {
	return PageBody(app.Div().Text("content goes here!!"))
}
