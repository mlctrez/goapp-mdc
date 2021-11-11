package selectm

import (
	_ "embed"
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

type MDCSelect struct {
	app.Compo
	base.JsUtil
	autoinit.AutoInit
	Id         string
	Width      int
	Items      list.Items
	Label      string
	ItemsLabel string
	Outlined   bool
	Required   bool
	api        app.Value
}

func (s *MDCSelect) Render() app.UI {
	selectDiv := app.Div().Class()
	selectDiv.Class("mdc-select").ID(s.Id)
	selectDiv.DataSet("mdc-auto-init", autoinit.MDCSelect)
	if s.Outlined {
		selectDiv.Class("mdc-select--outlined")
	} else {
		selectDiv.Class("mdc-select--filled")
	}
	if s.Required {
		selectDiv.Class("mdc-select--required")
	}

	if s.Width == 0 {
		// a reasonable default
		s.Width = 200
	}
	selectDiv.Style("width", fmt.Sprintf("%dpx", s.Width))

	labelId := fmt.Sprintf("%s-label", s.Id)
	selectedTextId := fmt.Sprintf("%s-selectedText", s.Id)

	var anchorBody []app.UI

	labelClass := "mdc-floating-label"
	if s.Items.SelectedItemText() != "" {
		labelClass += " mdc-floating-label--float-above"
	}

	selectedTextContainer := app.Span().Class("mdc-select__selected-text-container").Body(
		app.Span().ID(selectedTextId).Class("mdc-select__selected-text").Text(
			s.Items.SelectedItemText(),
		),
	)

	if s.Outlined {
		anchorBody = append(anchorBody,
			app.Span().Class("mdc-notched-outline").Body(
				app.Span().Class("mdc-notched-outline__leading"),
				app.Span().Class("mdc-notched-outline__notch").Body(
					app.Span().ID(labelId).Class(labelClass).Text(s.Label),
				),
				app.Span().Class("mdc-notched-outline__trailing"),
			),
			selectedTextContainer,
			dropdownIcon(),
		)
	} else {
		anchorBody = append(anchorBody,
			app.Span().Class("mdc-select_ripple"),
			app.Span().ID(labelId).Class(labelClass).Text(s.Label),
			selectedTextContainer,
			dropdownIcon(),
			app.Span().Class("mdc-line-ripple"),
		)
	}

	selectDiv.Body(
		app.Div().Class("mdc-select__anchor").Role("button").Aria("haspopup", "listbox").
			Aria("expanded", "false").Aria("required", s.Required).
			Aria("labeledby", labelId+" "+selectedTextId).Body(anchorBody...),
		s.Menu(),
	)

	return selectDiv
}

func dropdownIcon() app.HTMLSpan {
	return app.Span().Class("mdc-select__dropdown-icon").Body(app.Raw(DropDownIcon))
}

func (s *MDCSelect) Menu() app.UI {
	menu := app.Div().Class("mdc-select__menu")
	menu.Class("mdc-menu")
	menu.Class("mdc-menu-surface")
	menu.Class("mdc-menu-surface--fullwidth")
	menu.Body(
		app.Ul().Class("mdc-deprecated-list").Role("listbox").
			Aria("label", s.ItemsLabel).Body(s.Items.UIList()...),
	)

	return menu
}

var _ app.Mounter = (*MDCSelect)(nil)

type Event string

var MDCSelectChange Event = "MDCSelect:change"

func (s *MDCSelect) OnMount(ctx app.Context) {
	e := s.JSValue()
	s.MdcAutoInit(e)
	s.api = s.AutoInitComponent(e, autoinit.MDCSelect)
	e.Call("addEventListener", string(MDCSelectChange), app.FuncOf(s.event(ctx, MDCSelectChange)))
}

func (s *MDCSelect) event(ctx app.Context, change Event) func(this app.Value, args []app.Value) interface{} {
	return func(this app.Value, args []app.Value) interface{} {
		tags := app.Tags{}
		tags.Set("value", s.JsValueAt(args[0], "detail.value", false).String())
		tags.Set("index", s.JsValueAt(args[0], "detail.index", false).Int())
		ctx.NewActionWithValue(string(MDCSelectChange), s, tags)
		return nil
	}
}

//go:embed DropDownIcon.svg
var DropDownIcon string
