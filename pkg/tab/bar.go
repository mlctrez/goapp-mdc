// Package tab is the implementation of mdc-tab-bar.
//
// https://github.com/material-components/material-components-web/tree/master/packages/mdc-tab-bar
//
package tab

import (
	"log"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

// Bar is the mdc-tab-bar component containing the Tabs.
type Bar struct {
	app.Compo
	base.JsUtil
	autoinit.AutoInit
	Id   string
	Tabs Tabs
	api  app.Value
}

func (b *Bar) Render() app.UI {
	tabBar := app.Div().Class("mdc-tab-bar")
	API.DataMdcAutoInitDiv(tabBar)
	tabBar.Role("tablist")
	if b.Id != "" {
		tabBar.ID(b.Id)
	}
	scrollContent := app.Div().Class("mdc-tab-scroller__scroll-content")
	scrollContent.Body(b.Tabs.UIList()...)

	tabBar.Body(
		app.Div().Class("mdc-tab-scroller").Body(
			app.Div().Class("mdc-tab-scroller__scroll-area").Body(
				scrollContent,
			),
		),
	)
	return tabBar
}

func (b *Bar) OnMount(ctx app.Context) {
	e := b.JSValue()
	b.api = b.AutoInitComponent(e, API)
	e.Call("addEventListener", string(Activated), app.FuncOf(b.event(ctx, Activated)))
	ctx.Handle(string(Activate), b.handleActivate)
}

// API is the auto init data value for this component.
const API = autoinit.MDCTabBar

type EventType string

const Activated EventType = "MDCTabBar:activated"
const Activate EventType = "MDCTabBar:activate"

func (b *Bar) event(ctx app.Context, event EventType) func(this app.Value, args []app.Value) interface{} {

	return func(this app.Value, args []app.Value) interface{} {
		if len(args) > 0 {
			index := b.JsValueAt(args[0], "detail.index", true)
			if !index.IsUndefined() {
				ctx.NewActionWithValue(string(Activated), b, app.T("index", index.Int()))
			}
		}
		return nil
	}
}

func (b *Bar) handleActivate(ctx app.Context, action app.Action) {
	if action.Name == string(Activate) && action.Value == b {
		indexString := action.Tags.Get("index")
		atoi, err := strconv.Atoi(indexString)
		if err != nil {
			log.Println("error converting tab index", err)
			return
		}
		b.api.Call("activateTab", atoi)
	}
}
