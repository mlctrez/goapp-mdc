// Package autoinit allows interaction between go-app components and mdc.autoInit().
//
// See https://github.com/material-components/material-components-web/tree/master/packages/mdc-auto-init for full details
//
package autoinit

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type AutoInit struct{}

// TODO: more work to make sure that mdc is available when these are invoked.

// AutoInitComponent is short for MdcAutoInitWindow() and return name.GetFrom(compo).
func (ai *AutoInit) AutoInitComponent(compo app.Value, name MDCName) app.Value {
	ai.MdcAutoInitWindow()
	return name.GetFrom(compo)
}

// MdcAutoInit calls the js function mdc.autoInit() with the provided value
func (ai *AutoInit) MdcAutoInit(root app.Value) {
	app.Window().Get("mdc").Call("autoInit", root)
}

// MdcAutoInitWindow is short for MdcAutoInit(app.Window().get("Document"))
func (ai *AutoInit) MdcAutoInitWindow() {
	ai.MdcAutoInit(app.Window().Get("document"))
}



// GetFrom calls value.Get(string(MDCName))
func (n MDCName) GetFrom(value app.Value) app.Value {
	return value.Get(string(n))
}

