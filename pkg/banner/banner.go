// Package banner provides the banner component and interactions with that component.
//
// https://github.com/material-components/material-components-web/tree/master/packages/mdc-banner
package banner

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

// Banner component displays a prominent message and related optional actions.
//
// See https://github.com/material-components/material-components-web/tree/master/packages/mdc-banner#banner-example
// for usage of the Banner fields.
type Banner struct {
	app.Compo
	base.JsUtil
	autoinit.AutoInit
	Id       string
	Text     string
	Centered bool
	Fixed    bool
	Stacked  bool
	Buttons  []app.UI
	api      app.Value
}

// Render provides the markup to display the banner.
//
func (b *Banner) Render() app.UI {

	content := app.Div().Class("mdc-banner__content")
	content.Attr("role", "alertdialog").Aria("live", "assertlive")

	textWrapper := app.Div().Class("mdc-banner__graphic-text-wrapper").Body(
		app.Div().Class("mdc-banner__text").Text(b.Text),
	)
	actions := app.Div().Class("mdc-banner__actions").Body(b.Buttons...)

	content.Body(textWrapper, actions)
	if b.Fixed {
		content = app.Div().Class("mdc-banner__fixed").Body(content)
	}

	banner := app.Div().Class("mdc-banner")
	if b.Id != "" {
		banner.ID(b.Id)
	}
	API.DataMdcAutoInitDiv(banner)

	if b.Centered {
		banner.Class("mdc-banner--centered")
	}
	if b.Stacked {
		banner.Class("mdc-banner--mobile-stacked")
	}
	banner.Body(content)

	return banner
}

// API is the auto init data value for this component.
const API = autoinit.MDCBanner

// OnMount sets up the component's js api and the event handlers.
func (b *Banner) OnMount(ctx app.Context) {
	b.api = b.AutoInitComponent(b.JSValue(), API)
	e := b.JSValue()
	e.Call(base.AddEventListener, string(Opening), app.FuncOf(b.event(ctx, Opening)))
	e.Call(base.AddEventListener, string(Opened), app.FuncOf(b.event(ctx, Opened)))
	e.Call(base.AddEventListener, string(Closing), app.FuncOf(b.event(ctx, Closing)))
	e.Call(base.AddEventListener, string(Closed), app.FuncOf(b.event(ctx, Closed)))
	ctx.Handle(string(Open), b.handleOpenClose)
	ctx.Handle(string(Close), b.handleOpenClose)
}

// EventType defines the event strings that are sent to and from the banner control.
type EventType string

// Opening indicates when the banner begins its opening animation.
const Opening EventType = "MDCBanner:opening"

// Opened indicates when the banner finishes its opening animation.
const Opened EventType = "MDCBanner:opened"

// Closing indicates when the banner begins its closing animation.
const Closing EventType = "MDCBanner:closing"

// Closed indicates when the banner finishes its closing animation.
const Closed EventType = "MDCBanner:closed"

// Open is the event sent using ActionOpen to call the banner's api.open() method.
const Open EventType = "MDCBanner:open"

// Close is the event sent using OnClose to call the banner's api.close() method
const Close EventType = "MDCBanner:close"

func (b *Banner) event(ctx app.Context, event EventType) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		// determine what button was clicked if any from the event args
		reasonString := string(CloseUnspecified)
		if len(args) > 0 {
			detailReason := b.JsValueAt(args[0], "detail.reason", false)
			if !detailReason.IsUndefined() && detailReason.Type() == app.TypeNumber {
				switch detailReason.Int() {
				case ClosePrimary.Int():
					reasonString = string(ClosePrimary)
				case CloseSecondary.Int():
					reasonString = string(CloseSecondary)
				default:
					reasonString = string(CloseUnspecified)
				}
			}
		}

		ctx.NewActionWithValue(string(event), b, app.T(ReasonTag, reasonString))
		return nil
	}
}

func (b *Banner) handleOpenClose(ctx app.Context, action app.Action) {
	if b == action.Value {
		switch EventType(action.Name) {
		case Open:
			b.api.Call("open")
		case Close:
			atoi, err := strconv.Atoi(action.Tags.Get(ReasonTag))
			if err != nil {
				atoi = CloseUnspecified.Int()
			}
			b.api.Call("close", atoi)
		}
	}
}

// ActionOpen sends the Open event to the banner component.
func (b *Banner) ActionOpen(ctx app.Context) {
	ctx.NewActionWithValue(string(Open), b)
}

// ActionClose sends the Close event to the banner component with the provided reason.
func (b *Banner) ActionClose(ctx app.Context, reason CloseReason) {
	ctx.NewActionWithValue(string(Close), b, app.T(ReasonTag, reason.Int()))
}

// OnClose registers the onClose callback function to handle the Close event emitted by the banner js.
func (b *Banner) OnClose(ctx app.Context, onClose func(ctx app.Context, reason string)) {
	ctx.Handle(string(Closed), func(context app.Context, action app.Action) {
		if action.Value == b {
			onClose(context, action.Tags.Get(ReasonTag))
		}
	})
}

type CloseReason string
type CloseReasonInt int

const ClosePrimary CloseReason = "primary"
const CloseSecondary CloseReason = "secondary"
const CloseUnspecified CloseReason = "unspecified"

const ClosePrimaryInt CloseReasonInt = 0
const CloseSecondaryInt CloseReasonInt = 1
const CloseUnspecifiedInt CloseReasonInt = 3

const ReasonTag = "reason"

func (bcr CloseReason) Int() int {
	switch bcr {
	case ClosePrimary:
		return int(ClosePrimaryInt)
	case CloseSecondary:
		return int(CloseSecondaryInt)
	default:
		return int(CloseUnspecifiedInt)
	}
}
