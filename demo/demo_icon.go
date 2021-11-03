package demo

import (
	"fmt"
	"sort"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

type IconDemo struct {
	app.Compo
	base.JsUtil
	counterOne    *base.Counter
	toggleOne     *icon.Button
	toggleTwo     *icon.Button
	iconGroupList *list.List
}

func iconGroupNamesSorted() (result []string) {
	for s := range icon.AllGroupFunctions() {
		result = append(result, s)
	}
	sort.Strings(result)
	return
}

func (d *IconDemo) Render() app.UI {
	if d.counterOne == nil {
		d.counterOne = &base.Counter{Label: "bookmark"}
		d.toggleOne = &icon.Button{Id: d.UUID(), Icon: icon.MIFavorite,
			IconOff: icon.MIFavoriteBorder, AriaOn: "remove from favorites", AriaOff: "add to favorites"}
		d.toggleTwo = &icon.Button{Id: d.UUID(), Icon: icon.MIFavorite,
			IconOff: icon.MIFavoriteBorder, AriaOn: "remove from favorites", AriaOff: "add to favorites"}
		d.toggleOne.ButtonToggleChange = func(isOn bool) { d.toggleTwo.SetState(isOn) }
		d.toggleTwo.ButtonToggleChange = func(isOn bool) { d.toggleOne.SetState(isOn) }

		d.iconGroupList = &list.List{Type: list.SingleSelection, Id: "iconGroupList"}
		groups := list.Items{}
		for _, g := range iconGroupNamesSorted() {
			groups = append(groups, &list.Item{Text: g})
		}
		d.iconGroupList.Items = groups.UIList()

	}

	body := layout.Grid().Body(
		layout.Inner().Style("display", "flex").Body(
			layout.Cell().Body(
				&icon.Button{Id: d.UUID(), Icon: icon.MIBookmark,
					AriaOff: "bookmark this", Callback: d.IconButtonClicked}, d.counterOne),
			layout.Cell().Body(d.toggleOne, d.toggleTwo),
		),
		layout.Inner().Style("display", "flex").Body(
			layout.CellWide().Body(app.Text("Material Icon Groups")),
			layout.Cell().Body(d.iconGroupList),
		),
	)

	return PageBody(body)
}

func (d *IconDemo) IconButtonClicked(button app.HTMLButton) {
	button.OnClick(func(ctx app.Context, e app.Event) {
		fmt.Println("you clicked bookmark")
		d.counterOne.Count += 1
		d.counterOne.Update()
	})
}
