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
	Id   string
	List *list.List
	Open bool
	Type Type
	api  app.Value
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
	if d.Open {
		aside.Class("mdc-drawer--open")
	}

	switch d.Type {
	case Modal:
		aside.Class("mdc-drawer--modal")
		aside.DataSet("mdc-auto-init", "MDCDrawer")
	case Dismissible:
		aside.Class("mdc-drawer--dismissible")
		aside.DataSet("mdc-auto-init", "MDCDrawer")
	case Bottom:
		panic("MDC Web does not currently support bottom navigation drawers")
	case Standard:
		// no additional classes
	}
	return aside
}

type EventType string

const Open EventType = "MDCDrawer:open"
const Opened EventType = "MDCDrawer:opened"
const Close EventType = "MDCDrawer:close"
const Closed EventType = "MDCDrawer:closed"
const Toggle EventType = "MDCDrawer:toggle"

func (d *Drawer) OnMount(ctx app.Context) {
	switch d.Type {
	case Modal, Dismissible:
		e := d.JSValue()
		// TODO: can autoInit be called only on this element and document?
		app.Window().Get("mdc").Call("autoInit")
		d.api = e.Get("MDCDrawer")
		e.Call("addEventListener", string(Opened), app.FuncOf(d.event(ctx, Opened)))
		e.Call("addEventListener", string(Closed), app.FuncOf(d.event(ctx, Closed)))
		ctx.Handle(string(Open), d.handle)
		ctx.Handle(string(Close), d.handle)
		ctx.Handle(string(Toggle), d.handle)
	}
}

func (d *Drawer) event(ctx app.Context, event EventType) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		ctx.NewActionWithValue(string(event), d)
		return nil
	}
}

func (d *Drawer) handle(ctx app.Context, action app.Action) {
	if d == action.Value && d.api != nil {
		switch EventType(action.Name) {
		case Open:
			d.api.Set("open", true)
		case Close:
			d.api.Set("open", false)
		case Toggle:
			current := d.api.Get("open").Bool()
			d.api.Set("open", !current)
		}
	}
}

func (d *Drawer) ActionOpen(ctx app.Context) {
	ctx.NewActionWithValue(string(Open), d)
}

func (d *Drawer) ActionClose(ctx app.Context) {
	ctx.NewActionWithValue(string(Close), d)
}

func (d *Drawer) ActionToggle(ctx app.Context) {
	ctx.NewActionWithValue(string(Toggle), d)
}

func Scrim() app.UI {
	return app.Div().Class("mdc-drawer-scrim")
}
