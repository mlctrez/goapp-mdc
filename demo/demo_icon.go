package demo

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type IconDemo struct {
	app.Compo
	base.JsUtil
	counterOne *base.Counter
	toggleOne  *icon.Button
	toggleTwo  *icon.Button
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
	}

	body := layout.Grid().Body(
		layout.Inner().Body(
			layout.Cell().Body(
				&icon.Button{Id: d.UUID(), Icon: icon.MIBookmark,
					AriaOff: "bookmark this", Callback: d.IconButtonClicked}, d.counterOne),
			layout.Cell().Body(d.toggleOne, d.toggleTwo),
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
