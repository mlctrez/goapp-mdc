package list

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/radio"
)

type Item struct {
	app.Compo
	base.JsUtil
	Type      ItemType
	Graphic   icon.MaterialIcon
	Value     string
	Text      string
	Secondary string
	Href      string
	Name      string
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
			if i.Text != "" {
				content = append(content, app.Text(i.Text))
			}
		} else {
			content = append(content,
				app.Span().Class("mdc-deprecated-list-item__primary-text").
					Text(i.Text))
			content = append(content,
				app.Span().Class("mdc-deprecated-list-item__secondary-text").
					Text(i.Secondary))
		}
	case ItemTypeRadio:
		radId := i.id
		if radId == "" {
			radId = i.UUID()
		}
		content = append(content, app.Span().Class("mdc-deprecated-list-item__ripple"))
		content = append(content, app.Span().Class("mdc-deprecated-list-item__graphic").Body(
			&radio.Radio{Id: radId, Name: i.Name, Value: i.Value},
		))
		content = append(content, app.Label().Class("mdc-deprecated-list-item__text").For(radId).Text(i.Text))

	case ItemTypeCheckbox:
		cbId := i.id
		if cbId == "" {
			cbId = i.UUID()
		}
		content = append(content, app.Span().Class("mdc-deprecated-list-item__ripple"))
		content = append(content, app.Span().Class("mdc-deprecated-list-item__graphic").Body(
			&checkbox.Checkbox{Id: cbId},
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
		if i.Value != "" || i.Text == "" {
			root.DataSet("value", i.Value)
		}
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
			app.If(len(content) > 0,
				app.Span().Class("mdc-deprecated-list-item__text").Body(content...),
			).Else(),
		)
	}

	return root.UI()
}

func (i *Item) OnMount(_ app.Context) {
	i.MDCRippleVal(i.JSValue())
}
