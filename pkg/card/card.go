package card

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Card struct {
	app.Compo
	base.JsUtil
	Id            string
	Width         int
	Height        int
	Outlined      bool
	Padding       int
	PrimaryAction []app.UI
	ActionButtons []app.UI
	ActionIcons   []app.UI
}

func (c *Card) Render() app.UI {
	root := app.Div().ID(c.Id)
	root.Class("mdc-card")
	if c.Padding > 0 {
		root.Style("padding", fmt.Sprintf("%dpx", c.Padding))
	}
	if c.Width > 0 {
		root.Style("width", fmt.Sprintf("%dpx", c.Width))
	}
	if c.Height > 0 {
		root.Style("height", fmt.Sprintf("%dpx", c.Height))
	}
	// https://github.com/material-components/material-components-web/tree/master/packages/mdc-card#outlined-card
	if c.Outlined {
		root.Class("mdc-card--outlined")
	}

	var rootBody []app.UI

	// https://github.com/material-components/material-components-web/tree/master/packages/mdc-card#primary-action
	if len(c.PrimaryAction) > 0 {
		primaryAction := app.Div()
		primaryAction.Class("mdc-card__primary-action").TabIndex(0)
		var contentWithRipple = c.PrimaryAction
		contentWithRipple = append(contentWithRipple, app.Div().Class("mdc-card__ripple"))
		primaryAction.Body(contentWithRipple...)
		rootBody = append(rootBody, primaryAction)
	}

	if len(c.ActionButtons) > 0 || len(c.ActionIcons) > 0 {
		var actionsBody []app.UI
		if len(c.ActionButtons) > 0 {
			actionsBody = append(actionsBody, app.Div().Class("mdc-card__action-buttons").Body(c.ActionButtons...))
		}
		if len(c.ActionIcons) > 0 {
			actionsBody = append(actionsBody, app.Div().Class("mdc-card__action-icons").Body(c.ActionIcons...))
		}
		actions := app.Div().Class("mdc-card__actions").Body(actionsBody...)
		rootBody = append(rootBody, actions)
	}
	root.Body(rootBody...)
	return root
}

func (c *Card) OnMount(_ app.Context) {
	c.MDCRipple(c.Id)
}
