package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
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
