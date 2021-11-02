package demo

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

type ListDemo struct {
	app.Compo
	base.JsUtil
}

func (d *ListDemo) Render() app.UI {

	regularList := list.Items{&list.Item{Text: "item one"}, &list.Item{Text: "item two"}, &list.Item{Text: "item three"}}.Select(-1)
	twoLineList := list.Items{
		&list.Item{Text: "item one", Secondary: "item one subtext"},
		&list.Item{Text: "item two", Secondary: "item two subtext"},
		&list.Item{Text: "item three", Secondary: "item three subtext"}}.Select(-1)

	groupedListOne := list.Items{&list.Item{Text: "group 1-1"}, &list.Item{Text: "group 2-2"}}.Select(0)
	groupedListTwo := list.Items{&list.Item{Text: "group 2-1"}, &list.Item{Text: "group 2-2"}}.Select(1)

	singleSelectionList := list.Items{&list.Item{Text: "item one"}, &list.Item{Text: "item two"},
		&list.Item{Text: "item three"}, &list.Item{Text: "item four"}}.Select(2)

	dividedList := list.Items{&list.Item{Text: "item one"}, &list.Item{Text: "item two before divider"},
		&list.Item{Type: list.ItemTypeDivider}, &list.Item{Text: "item three after divider"}, &list.Item{Text: "item four"}}
	dividedList.Select(0)

	// TODO: build out radio button component first
	//radioGroupList := Items{&Item{Text: "item one"}, &Item{Text: "item two"},
	//	&Item{Text: "item three"}, &Item{Text: "item four"}}.Select(2)

	checkboxGroupList := make(list.Items, 4)
	for i := range checkboxGroupList {
		checkboxGroupList[i] = &list.Item{Type: list.ItemTypeCheckbox, Text: fmt.Sprintf("checkbox %d", i)}
	}
	checkboxGroupList.Select(-1)

	body := layout.Grid().Body(layout.Inner().Body(
		layout.Cell().Body(
			app.P().Text("regular"), &list.List{Id: "regularList", Items: regularList.UIList()}),
		layout.Cell().Body(
			app.P().Text("two line"), &list.List{Id: "twoLineList", TwoLine: true, Items: twoLineList.UIList()}),
		layout.Cell().Body(
			app.P().Text("grouped"),
			&list.Group{Items: []*list.GroupItem{
				{SubHeader: "group 1", List: &list.List{Id: "groupedList1", Items: groupedListOne.UIList()}},
				{SubHeader: "group 2", List: &list.List{Id: "groupedList2", Items: groupedListTwo.UIList()}},
			}},
		),
		layout.Cell().Body(app.P().Text("divided List"), &list.List{Id: "dividedList", Items: dividedList.UIList()}),
		layout.Cell().Body(
			app.P().Text("single select"),
			&list.List{Id: "singleSelectionList", Type: list.SingleSelection, Items: singleSelectionList.UIList()},
		),
		//layout.Cell().Body(
		//	app.P().Text("radio group"),
		//	&List{Id: "radioGroupList", Type: RadioGroup, Items: radioGroupList.UIList()},
		//),
		layout.Cell().Body(
			app.P().Text("checkbox group"),
			&list.List{Id: "checkboxGroupList", Type: list.CheckBox, Items: checkboxGroupList.UIList()},
		),
	))
	return PageBody(body)

}
