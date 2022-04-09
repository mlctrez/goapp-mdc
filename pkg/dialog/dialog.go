package dialog

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Dialog struct {
	app.Compo
	base.JsUtil
	autoinit.AutoInit
	Id      string
	Title   []app.UI
	Content []app.UI
	Buttons []app.UI
	api     app.Value
}

func (d *Dialog) Render() app.UI {
	dialog := app.Div().Class("mdc-dialog").ID(d.Id)
	API.DataMdcAutoInitDiv(dialog)

	dialogSurface := app.Div().Class("mdc-dialog__surface").Attr("role", "alertdialog")
	dialogSurface.Aria("modal", "true")
	if d.Title != nil {
		dialogSurface.Aria("labelledby", d.Id+"-title")
	}
	if d.Content != nil {
		dialogSurface.Aria("describedby", d.Id+"-content")
	}
	dialogSurface.Body(
		app.If(d.Title != nil, app.Div().Class("mdc-dialog__title").ID(d.Id+"-title").Body(d.Title...)).Else(),
		app.If(d.Content != nil, app.Div().Class("mdc-dialog__content").ID(d.Id+"-content").Body(d.Content...)).Else(),
		app.Div().Class("mdc-dialog__actions").Body(d.Buttons...),
	)

	dialog.Body(
		app.Div().Class("mdc-dialog__container").Body(dialogSurface),
		app.Div().Class("mdc-dialog__scrim"),
	)
	return dialog
}

// EventType defines the event strings that are sent to and from the dialog control.
type EventType string

// Opening indicates when the dialog begins its opening animation.
const Opening EventType = "MDCDialog:opening"

// Opened indicates when the dialog finishes its opening animation.
const Opened EventType = "MDCDialog:opened"

// Closing indicates when the dialog begins its closing animation.
const Closing EventType = "MDCDialog:closing"

// Closed indicates when the dialog finishes its closing animation.
const Closed EventType = "MDCDialog:closed"

// Open is the event sent using ActionOpen to call the dialog's api.open() method.
const Open EventType = "MDCDialog:open"

// Close is the event sent using ActionClose to call the dialog's api.close() method
const Close EventType = "MDCDialog:close"

// API is the auto init data value for this component.
const API = autoinit.MDCDialog

const ActionTag = "action"

func (d *Dialog) OnMount(ctx app.Context) {
	d.api = d.AutoInitComponent(d.JSValue(), API)
	e := d.JSValue()
	e.Call(base.AddEventListener, string(Opening), app.FuncOf(d.event(ctx, Opening)))
	e.Call(base.AddEventListener, string(Opened), app.FuncOf(d.event(ctx, Opened)))
	e.Call(base.AddEventListener, string(Closing), app.FuncOf(d.event(ctx, Closing)))
	e.Call(base.AddEventListener, string(Closed), app.FuncOf(d.event(ctx, Closed)))
	ctx.Handle(string(Open), d.handleOpenClose)
	ctx.Handle(string(Close), d.handleOpenClose)
}

func (d *Dialog) handleOpenClose(ctx app.Context, action app.Action) {
	if d == action.Value || d.Id == action.Value {
		switch EventType(action.Name) {
		case Open:
			d.api.Call("open")
		case Close:
			d.api.Call("close", action.Tags.Get(ActionTag))
		}
	}
}

// ActionOpen sends the Open event to the dialog component.
func (d *Dialog) ActionOpen(ctx app.Context) {
	ctx.NewActionWithValue(string(Open), d)
}

// ActionClose sends the Close event to the dialog component with the provided action.
func (d *Dialog) ActionClose(ctx app.Context, action string) {
	ctx.NewActionWithValue(string(Close), d, app.T(ActionTag, action))
}

func (d *Dialog) event(ctx app.Context, event EventType) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		actionString := ""
		if len(args) > 0 {
			detailAction := d.JsValueAt(args[0], "detail.action", false)
			if !detailAction.IsUndefined() && detailAction.Type() == app.TypeString {
				actionString = detailAction.String()
			}
		}

		ctx.NewActionWithValue(string(event), d, app.T(ActionTag, actionString))
		return nil
	}
}

// Open opens this dialog
// Deprecated: Use ActionOpen or context.NewActionWithValue using the id of the dialog
func (d *Dialog) Open() {
	if !d.api.IsUndefined() {
		d.api.Call("open")
	}
}
