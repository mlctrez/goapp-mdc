package icon

import (
	_ "embed"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

//go:embed demo.go
var code string

type Code struct {
	app.Compo
}

func (c *Code) Render() app.UI {
	return app.Pre().Body(app.Code().Class("language-go").Text(code))
}

func (c *Code) OnMount(ctx app.Context) {
	ctx.Defer(c.highlightCode)
}

func (c *Code) OnUpdate(ctx app.Context) {
	ctx.Defer(c.highlightCode)
}

func (c *Code) highlightCode(ctx app.Context) {
	app.Window().Get("Prism").Call("highlightAll")
}
