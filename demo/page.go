package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

// PageBody applies the navigation, update banner, and demo page layout to the provided pageContent.
func PageBody(pageContent ...app.UI) app.UI {

	content := []app.UI{&Navigation{}}
	content = append(content, pageContent...)

	return app.Div().Body(
		&AppUpdateBanner{},

		app.Div().Style("display", "flex").Body(content...),
	)
}

func FlexGrid(cells ...app.UI) app.UI {
	return layout.Grid().Body(
		layout.Inner().Style("display", "flex").Body(cells...),
	)
}
