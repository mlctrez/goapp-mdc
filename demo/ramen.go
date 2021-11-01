package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Ramen struct {
	app.Compo
	base.JsUtil
}

func (r *Ramen) Render() app.UI {
	return PageBody(app.Div().Text("ramen content goes here!!"))
}
