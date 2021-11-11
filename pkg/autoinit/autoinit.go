package autoinit

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type AutoInit struct{}

// TODO: validate that mdc javascript has initialized

// AutoInitComponent is short for MdcAutoInitWindow() and compo.get(name) .
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

type MDCName string

// GetFrom calls value.Get(string(MDCName))
func (n MDCName) GetFrom(value app.Value) app.Value {
	return value.Get(string(n))
}

const MDCBanner MDCName = "MDCBanner"
const MDCCheckbox MDCName = "MDCCheckbox"
const MDCChip MDCName = "MDCChip"
const MDCChipSet MDCName = "MDCChipSet"
const MDCCircularProgress MDCName = "MDCCircularProgress"
const MDCDataTable MDCName = "MDCDataTable"
const MDCDialog MDCName = "MDCDialog"
const MDCDrawer MDCName = "MDCDrawer"
const MDCFloatingLabel MDCName = "MDCFloatingLabel"
const MDCFormField MDCName = "MDCFormField"
const MDCIconButtonToggle MDCName = "MDCIconButtonToggle"
const MDCLineRipple MDCName = "MDCLineRipple"
const MDCLinearProgress MDCName = "MDCLinearProgress"
const MDCList MDCName = "MDCList"
const MDCMenu MDCName = "MDCMenu"
const MDCMenuSurface MDCName = "MDCMenuSurface"
const MDCNotchedOutline MDCName = "MDCNotchedOutline"
const MDCRadio MDCName = "MDCRadio"
const MDCRipple MDCName = "MDCRipple"
const MDCSegmentedButton MDCName = "MDCSegmentedButton"
const MDCSelect MDCName = "MDCSelect"
const MDCSlider MDCName = "MDCSlider"
const MDCSnackbar MDCName = "MDCSnackbar"
const MDCSwitch MDCName = "MDCSwitch"
const MDCTabBar MDCName = "MDCTabBar"
const MDCTextField MDCName = "MDCTextField"
const MDCTooltip MDCName = "MDCTooltip"
const MDCTopAppBar MDCName = "MDCTopAppBar"
