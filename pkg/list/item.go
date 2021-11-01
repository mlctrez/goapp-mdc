package list

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
)

type Item struct {
	app.Compo
	base.JsUtil
	Type      ItemType
	Graphic   icon.MaterialIcon
	Text      string
	Secondary string
	Href      string
	state     ItemSelectState
	// for testing only
	id string
}

type ItemType uint

const (
	ItemTypeNone ItemType = iota
	ItemTypeOption
	ItemTypeRadio
	ItemTypeCheckbox
	ItemTypeDivider
	ItemTypeAnchor
)

func (i *Item) Render() app.UI {

	var content []app.UI

	switch i.Type {
	case ItemTypeDivider:
		return app.Li().Attr("role", "separator").Class("mdc-deprecated-list-divider")
	case ItemTypeNone, ItemTypeOption, ItemTypeAnchor:
		if i.Secondary == "" {
			content = append(content, app.Text(i.Text))
		} else {
			content = append(content,
				app.Span().Class("mdc-deprecated-list-item__primary-text").
					Text(i.Text))
			content = append(content,
				app.Span().Class("mdc-deprecated-list-item__secondary-text").
					Text(i.Secondary))
		}
	case ItemTypeRadio:
		panic("not implemented")
	case ItemTypeCheckbox:
		cbId := i.id
		if cbId == "" {
			cbId = i.UUID()
		}
		// TODO: some refactoring of the checkbox package to allow use here?
		// Maybe not since the ripple span is located differently
		content = append(content, app.Span().Class("mdc-deprecated-list-item__ripple"))
		content = append(content, app.Span().Class("mdc-deprecated-list-item__graphic").Body(
			app.Div().Class("mdc-checkbox").Body(
				app.Input().Type("checkbox").Class("mdc-checkbox__native-control").ID(cbId),
				checkbox.MDCCheckboxBackground(),
			),
		))
		content = append(content, app.Label().Class("mdc-deprecated-list-item__text").For(cbId).Text(i.Text))
	}

	root := i.adapt()

	if i.Href != "" {
		root.Href(i.Href)
	}

	root.Class("mdc-deprecated-list-item")

	switch i.state {
	case ItemSelectStateSelected:
		root.Class("mdc-deprecated-list-item--selected")
		root.Aria("selected", "true")
		root.TabIndex(0)
	case ItemSelectStateTabZero:
		root.TabIndex(0)
	case ItemSelectStateNotSelected:
		root.Aria("selected", "false")
	}

	switch i.Type {
	case ItemTypeOption:
		root.Attr("role", "option")
	case ItemTypeRadio:
		root.Attr("role", "radio")
	case ItemTypeCheckbox:
		root.Attr("role", "checkbox")
	}

	switch i.Type {
	case ItemTypeCheckbox:
		root.Body(content...)
	default:
		root.Body(
			app.Span().Class("mdc-deprecated-list-item__ripple"),
			app.If(i.Graphic != "", i.Graphic.IItemGraphic()).Else(),
			app.Span().Class("mdc-deprecated-list-item__text").Body(content...),
		)
	}

	return root.UI()
}

func (i *Item) OnMount(_ app.Context) {
	i.MDCRippleVal(i.JSValue())
}