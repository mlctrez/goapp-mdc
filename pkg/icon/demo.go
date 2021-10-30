package icon

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type Demo struct {
	app.Compo
	base.JsUtil
	counterOne *base.Counter
	toggleOne  *Button
	toggleTwo  *Button
}

func (d *Demo) Render() app.UI {
	if d.counterOne == nil {
		d.counterOne = &base.Counter{Label: "bookmark"}
		d.toggleOne = &Button{Id: d.UUID(), Icon: MIFavorite, IconOff: MIFavoriteBorder, AriaOn: "remove from favorites", AriaOff: "add to favorites"}
		d.toggleTwo = &Button{Id: d.UUID(), Icon: MIFavorite, IconOff: MIFavoriteBorder, AriaOn: "remove from favorites", AriaOff: "add to favorites"}
		d.toggleOne.ButtonToggleChange = func(isOn bool) { d.toggleTwo.SetState(isOn) }
		d.toggleTwo.ButtonToggleChange = func(isOn bool) { d.toggleOne.SetState(isOn) }
	}
	return app.Div().Body(layout.Grid().Body(layout.Inner().Body(
		layout.Cell().Body(
			&Button{Id: d.UUID(), Icon: MIBookmark, AriaOff: "bookmark this", Callback: d.IconButtonClicked}, d.counterOne,
		),
		layout.Cell().Body(d.toggleOne, d.toggleTwo),
	),
	), app.Hr(), &Code{})
}

func (d *Demo) IconButtonClicked(button app.HTMLButton) {
	button.OnClick(func(ctx app.Context, e app.Event) {
		fmt.Println("you clicked bookmark")
		d.counterOne.Count += 1
		d.counterOne.Update()
	})
}
