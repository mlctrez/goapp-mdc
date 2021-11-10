package snackbar

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type SnackBar struct {
	app.Compo
	base.JsUtil
	autoinit.AutoInit
	Leading          bool
	Stacked          bool
	LabelText        string
	ActionButtonText string
	api              app.Value
}

// https://github.com/material-components/material-components-web/tree/master/packages/mdc-snackbar#javascript-api

type Event string

const Opening Event = "MDCSnackbar:opening"
const Opened Event = "MDCSnackbar:opened"
const Closing Event = "MDCSnackbar:closing"
const Closed Event = "MDCSnackbar:closed"

const Open Event = "MDCSnackbar:open"
const Close Event = "MDCSnackbar:close"
const LabelText Event = "MDCSnackbar:labelText"
const ActionButtonText Event = "MDCSnackbar:actionButtonText"

func (s *SnackBar) IsOpen() bool {
	return s.api.Get("isOpen").Bool()
}

const TimeoutMsField = "timeoutMs"

// TimeoutMs sets the value of MDCSnackbar.timeoutMs and returns the current value.
//
// Per the api documentation, the following ranges are valid. When outside the range,
// the api value is not set, but the current value is still returned. i.e. calling with 0
// can be used for reading the current value.
//   -1     no timeout
//   4000   minimum ms value
//   10000  max ms value
func (s *SnackBar) TimeoutMs(ms int) app.Value {
	if ms == -1 || (ms >= 4000 && ms <= 10000) {
		s.api.Set(TimeoutMsField, ms)
	}
	return s.api.Get(TimeoutMsField)
}

var _ app.Mounter = (*SnackBar)(nil)

func (s *SnackBar) OnMount(ctx app.Context) {
	e := s.JSValue()
	s.MdcAutoInit(e)
	s.api = autoinit.MDCSnackbar.GetFrom(e)
	for _, evt := range []Event{Opening, Opened, Closing, Closed} {
		e.Call("addEventListener", string(evt), app.FuncOf(s.event(ctx, evt)))
	}
	for _, evt := range []Event{Open, Close} {
		ctx.Handle(string(evt), s.handle)
	}
}

func (s *SnackBar) event(ctx app.Context, evt Event) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		tags := app.Tags{}
		reason := s.JsValueAt(args[0], "detail.reason", false)
		if !reason.IsUndefined() {
			tags.Set("reason", reason)
		}
		ctx.NewActionWithValue(string(evt), s, tags)
		return nil
	}
}

func (s *SnackBar) handle(context app.Context, action app.Action) {
	if action.Value == s {
		switch action.Name {
		case string(Open):
			s.api.Call("open")
		case string(Close):
			switch action.Tags.Get("reason") {
			case "":
				s.api.Call("close")
			default:
				s.api.Call("close", action.Tags.Get("reason"))
			}

		}
	}
}

func (s *SnackBar) Render() app.UI {
	aside := app.Aside().Class("mdc-snackbar")
	aside.DataSet("mdc-auto-init", "MDCSnackbar")
	if s.Leading {
		aside.Class("mdc-snackbar--leading")
	}
	if s.Stacked {
		aside.Class("mdc-snackbar--stacked")
	}
	aside.Body(
		surface().Body(
			label().Text(s.LabelText),
			actions().Body(button(s.ActionButtonText)),
		),
	)
	return aside
}

func button(text string) app.HTMLButton {
	return app.Button().Type("button").Class("mdc-button mdc-snackbar__action").Body(
		app.Div().Class("mdc-button__ripple"),
		app.Span().Class("mdc-button__label").Text(text),
	)
}

func surface() app.HTMLDiv {
	return app.Div().Class("mdc-snackbar__surface").Role("status").Aria("relevant", "additions")
}

func label() app.HTMLDiv {
	return app.Div().Class("mdc-snackbar__label").Aria("atomic", "false")
}

func actions() app.HTMLDiv {
	return app.Div().Class("mdc-snackbar__actions").Aria("atomic", "true")
}
