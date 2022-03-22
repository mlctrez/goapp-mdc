// Package autoinit allows interaction between go-app components and mdc.autoInit().
//
// See https://github.com/material-components/material-components-web/tree/master/packages/mdc-auto-init
// for more detail.
//
// Typical use occur in OnMount() where the component element is initialized and the
// value attached to the component is referenced in a field of the go-app component.
//
//	type TopAppBar struct {
//		app.Compo
//		autoinit.AutoInit
//		.....
//	}
//
//	func (c *TopAppBar) OnMount(_ app.Context) {
//		c.api = c.AutoInitComponent(c.JSValue(), autoinit.MDCTopAppBar)
//	}
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

// DataMdcAutoInitDiv is a convenience method for setting the data-mdc-auto-init attribute on a Div.
func (n MDCName) DataMdcAutoInitDiv(div app.HTMLDiv) app.HTMLDiv {
	return div.DataSet(MdcAutoInitData, n)
}

// DataMdcAutoInitButton is a convenience method for setting the data-mdc-auto-init attribute on a Button.
func (n MDCName) DataMdcAutoInitButton(button app.HTMLButton) app.HTMLButton {
	return button.DataSet(MdcAutoInitData, n)
}

// DataMdcAutoInitHeader is a convenience method for setting the data-mdc-auto-init attribute on a Header.
func (n MDCName) DataMdcAutoInitHeader(header app.HTMLHeader) app.HTMLHeader {
	return header.DataSet(MdcAutoInitData, n)
}

// DataMdcAutoInitSpan is a convenience method for setting the data-mdc-auto-init attribute on a Span.
func (n MDCName) DataMdcAutoInitSpan(set app.HTMLSpan) app.HTMLSpan {
	return set.DataSet(MdcAutoInitData, n)
}


// MdcAutoInitData is the data attribute constant.
const MdcAutoInitData = "mdc-auto-init"
