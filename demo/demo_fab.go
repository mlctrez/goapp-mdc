package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/fab"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type FabDemo struct {
	app.Compo
	activeIndex int
}

func (d *FabDemo) Render() app.UI {

	return PageBody(FlexGrid(
		layout.Grid().Body(
			&fab.Fab{Id: id(), Icon: "favorite"},
			app.Text("regular"),
		),
		layout.Grid().Body(
			&fab.Fab{Id: id(), Icon: "favorite", Mini: true},
			app.Text("mini"),
		),
		layout.Grid().Body(
			&fab.Fab{Id: id(), Icon: "favorite", Extended: true, Label: "Favorite"},
			app.Text("extended"),
		),
	))

}
