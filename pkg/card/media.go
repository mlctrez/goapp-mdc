package card

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

// Media represents a div of class mdc-card__media to be placed in the primary action section of a card.
type Media struct {
	app.Compo
	base.JsUtil
	Title  string
	Width  int
	Height int
	Image  string
}

func (c *Media) Render() app.UI {

	width := fmt.Sprintf("%dpx", c.Width)
	height := fmt.Sprintf("%dpx", c.Height)

	root := app.Div()
	root.Class("mdc-card__media")
	root.Style("width", width)
	root.Style("height", height)
	root.Style("background-image", fmt.Sprintf("url('%s')", c.Image))
	root.Style("background-size", fmt.Sprintf("%s %s", width, height))
	if c.Title != "" {
		root.Body(app.Div().Class("mdc-card__media-content").Text(c.Title))
	}
	return root
}
