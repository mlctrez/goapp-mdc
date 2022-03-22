// Package bar implements the mdc-top-app-bar component.
//
// https://github.com/material-components/material-components-web/tree/master/packages/mdc-top-app-bar
//
package bar

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

// TopAppBar is the go-app component for mdc-top-app-bar.
type TopAppBar struct {
	app.Compo
	autoinit.AutoInit
	Title        string
	Fixed        bool
	Navigation   []app.HTMLButton
	Actions      []app.HTMLButton
	ScrollTarget string
	api          app.Value
}

// Render returns the mdc-top-app-bar header configured with the required classes and child elements.
func (c *TopAppBar) Render() app.UI {
	header := app.Header().Class("mdc-top-app-bar")
	API.DataMdcAutoInitHeader(header)

	if c.Fixed {
		header.Class("mdc-top-app-bar--fixed")
	}

	navBody := base.HTMLButtonUpdate(c.Navigation, func(button app.HTMLButton) {
		button.Class("mdc-top-app-bar__navigation-icon")
	})
	if c.Title != "" {
		navBody = append(navBody, app.Span().Class("mdc-top-app-bar__title").Text(c.Title))
	}
	actionBody := base.HTMLButtonUpdate(c.Actions, func(button app.HTMLButton) {
		button.Class("mdc-top-app-bar__action-item")
	})

	navClass := "mdc-top-app-bar__section mdc-top-app-bar__section--align-start"
	actionClass := "mdc-top-app-bar__section mdc-top-app-bar__section--align-end"
	header.Body(
		app.Div().Class("mdc-top-app-bar__row").Body(
			app.Section().Class(navClass).Body(navBody...),
			app.Section().Class(actionClass).Body(actionBody...),
		),
	)
	return header
}

// Main returns the main element with class for this TopAppBar
func (c *TopAppBar) Main() app.HTMLMain {
	return app.Main().Class("mdc-top-app-bar--fixed-adjust")
}

// API is the auto init data value for this component.
const API = autoinit.MDCTopAppBar

// OnMount sets up the component's js api and the scroll target if present.
func (c *TopAppBar) OnMount(_ app.Context) {
	c.api = c.AutoInitComponent(c.JSValue(), API)
	if c.ScrollTarget != "" {
		c.api.Call("setScrollTarget", app.Window().GetElementByID(c.ScrollTarget))
	}
}
