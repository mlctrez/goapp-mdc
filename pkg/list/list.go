package list

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type List struct {
	app.Compo
	base.JsUtil
	Id      string
	Items   []app.UI
	TwoLine bool
	Type    Type
	jsApi   app.Value
}

// Type drives the rendering of the list
type Type int

const (
	None Type = iota
	SingleSelection
	RadioGroup
	CheckBox
	Navigation
)

func (l *List) Render() app.UI {
	root := l.adapt()
	root.Class("mdc-deprecated-list")
	if l.Id != "" {
		root.ID(l.Id)
	}
	if l.TwoLine {
		root.Class("mdc-deprecated-list--two-line")
	}
	switch l.Type {
	case SingleSelection:
		root.Attr("role", "listbox")
	case RadioGroup:
		root.Attr("role", "radiogroup")
	case CheckBox:
		root.Attr("role", "group")
	}
	root.Body(l.Items...)
	return root.UI()
}

type EventType string

const Action EventType = "MDCList:action"
const Select EventType = "MDCList:select"

func (l *List) OnMount(ctx app.Context) {
	e := l.JSValue()
	l.jsApi = l.JsNewAtPath("mdc.list.MDCList", e)
	if l.Type == SingleSelection && l.jsApi.Truthy() {
		l.jsApi.Set("singleSelection", true)
	}
	e.Call("addEventListener", string(Action), app.FuncOf(l.event(ctx, Action)))
}

func (l *List) Select(idx int) {
	for _, item := range l.Items {
		value := item.JSValue()
		if value.Truthy() {
			value.Call("blur")
		}
	}
	l.jsApi.Set("selectedIndex", idx)

}

func (l *List) event(ctx app.Context, action EventType) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		if len(args) < 1 {
			return nil
		}
		idx := l.JsValueAt(args[0], "detail.index", false)
		if !idx.IsUndefined() {
			ctx.NewActionWithValue(string(Select), l, app.T("index", idx.Int()))
		}
		return nil
	}
}
