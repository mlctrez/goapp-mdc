package drawer

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

// https://github.com/material-components/material-components-web/tree/master/packages/mdc-drawer#types
// There are three types of navigation drawers: standard (1), modal (2), and bottom (3).
// MDC Web does not currently support bottom navigation drawers.

type Drawer struct {
	app.Compo
	base.JsUtil
	Id    string
	List  *list.List
	Type  Type
	jsApi app.Value
}

type Type int

const (
	Standard Type = iota
	Modal
	Dismissible
	Bottom
)

func (d *Drawer) Render() app.UI {
	content := app.Div().Class("mdc-drawer__content").Body(d.List)
	aside := app.Aside().ID(d.Id).Class("mdc-drawer").Body(content)
	switch d.Type {
	case Modal:
		aside.Class("mdc-drawer--modal")
	case Dismissible:
		aside.Class("mdc-drawer--dismissible")
	case Bottom:
		panic("MDC Web does not currently support bottom navigation drawers")
	case Standard:
		// no additional classes
	}
	return aside
}

func (d *Drawer) OnMount(ctx app.Context) {
	e := app.Window().GetElementByID(d.Id)
	switch d.Type {
	case Dismissible, Modal:
		d.jsApi = d.JsNewAtPath("mdc.drawer.MDCDrawer", e)
	}
}

func Scrim() app.UI {
	return app.Div().Class("mdc-drawer-scrim")
}

//
//type NavLink struct {
//	Icon   icon.MaterialIcon
//	Active bool
//	Href   string
//	Text   string
//}
