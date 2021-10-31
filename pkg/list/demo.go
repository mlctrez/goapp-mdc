package list

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type Demo struct {
	app.Compo
	base.JsUtil
}

func (d *Demo) Render() app.UI {

	regularList := Items{&Item{Text: "item one"}, &Item{Text: "item two"}, &Item{Text: "item three"}}.Select(-1)
	twoLineList := Items{
		&Item{Text: "item one", Secondary: "item one subtext"},
		&Item{Text: "item two", Secondary: "item two subtext"},
		&Item{Text: "item three", Secondary: "item three subtext"}}.Select(-1)

	groupedListOne := Items{&Item{Text: "group 1-1"}, &Item{Text: "group 2-2"}}.Select(0)
	groupedListTwo := Items{&Item{Text: "group 2-1"}, &Item{Text: "group 2-2"}}.Select(1)

	singleSelectionList := Items{&Item{Text: "item one"}, &Item{Text: "item two"},
		&Item{Text: "item three"}, &Item{Text: "item four"}}.Select(2)

	dividedList := Items{&Item{Text: "item one"}, &Item{Text: "item two before divider"},
		&Item{Type: ItemTypeDivider}, &Item{Text: "item three after divider"}, &Item{Text: "item four"}}
	dividedList.Select(0)

	// TODO: build out radio button component first
	//radioGroupList := Items{&Item{Text: "item one"}, &Item{Text: "item two"},
	//	&Item{Text: "item three"}, &Item{Text: "item four"}}.Select(2)

	checkboxGroupList := make(Items, 4)
	for i := range checkboxGroupList {
		checkboxGroupList[i] = &Item{Type: ItemTypeCheckbox, Text: fmt.Sprintf("checkbox %d", i)}
	}
	checkboxGroupList.Select(-1)

	return layout.Grid().Body(layout.Inner().Body(
		layout.Cell().Body(
			app.P().Text("regular list"), &List{Id: "regularList", Items: regularList.UIList()}),
		layout.Cell().Body(
			app.P().Text("two line list"), &List{Id: "twoLineList", TwoLine: true, Items: twoLineList.UIList()}),
		layout.Cell().Body(
			app.P().Text("grouped List"),
			&Group{Items: []*GroupItem{
				{SubHeader: "group 1", List: &List{Id: "groupedList1", Items: groupedListOne.UIList()}},
				{SubHeader: "group 2", List: &List{Id: "groupedList2", Items: groupedListTwo.UIList()}},
			}},
		),
		layout.Cell().Body(app.P().Text("divided List"), &List{Id: "dividedList", Items: dividedList.UIList()}),
		layout.Cell().Body(
			app.P().Text("single selection"),
			&List{Id: "singleSelectionList", Type: SingleSelection, Items: singleSelectionList.UIList()},
		),
		//layout.Cell().Body(
		//	app.P().Text("radio group"),
		//	&List{Id: "radioGroupList", Type: RadioGroup, Items: radioGroupList.UIList()},
		//),
		layout.Cell().Body(
			app.P().Text("checkbox group"),
			&List{Id: "checkboxGroupList", Type: CheckBox, Items: checkboxGroupList.UIList()},
		),
	))

}
