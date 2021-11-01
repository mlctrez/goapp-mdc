package demo

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func PageBody(elems ...app.UI) app.UI {

	content := []app.UI{&Navigation{}}
	content = append(content, elems...)

	return app.Div().Body(&AppUpdateBanner{}, app.Div().Style("display","flex").Body(content...))
}
