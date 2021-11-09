package bar

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type TopAppBar struct {
	app.Compo
	Title      string
	Fixed      bool
	Navigation []app.HTMLButton
	Actions    []app.HTMLButton
	api        app.Value
}

var _ app.Mounter = (*TopAppBar)(nil)

func (c *TopAppBar) Render() app.UI {
	header := app.Header().Class("mdc-top-app-bar")
	if c.Fixed {
		header.Class("mdc-top-app-bar--fixed")
	}

	header.DataSet("mdc-auto-init", "MDCTopAppBar")

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

func (c *TopAppBar) OnMount(_ app.Context) {
	app.Window().Get("mdc").Call("autoInit", c.JSValue())
	c.api = c.JSValue().Get("MDCTopAppBar")
}
