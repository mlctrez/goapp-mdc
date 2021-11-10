package demo

import (
	"fmt"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/bar"
	"github.com/mlctrez/goapp-mdc/pkg/drawer"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

// PageBody applies the navigation, update banner, and demo page layout to the provided pageContent.
func PageBody(pageContent ...app.UI) app.UI {
	nav := &Navigation{Type: drawer.Dismissible}
	topBar := &bar.TopAppBar{Title: "go-app mdc", Fixed: false}

	topBar.Navigation = []app.HTMLButton{icon.MIMenu.Button().OnClick(func(ctx app.Context, e app.Event) {
		nav.drawer.ActionToggle(ctx)
	})}

	reloadButton := icon.MIRefresh.Button().OnClick(func(ctx app.Context, e app.Event) {
		ctx.Reload()
	})

	codeButton := icon.MICode.Button().OnClick(func(ctx app.Context, e app.Event) {
		fragment := strings.TrimPrefix(ctx.Page().URL().Path, "/") + ".go"
		ctx.Navigate(fmt.Sprintf("/code#%s", fragment))
	})

	githubButton := app.Button().Class(icon.MaterialIconsClass, icon.MaterialIconButton).
		Body(app.Raw(GitHubSvg)).OnClick(func(ctx app.Context, e app.Event) {
		app.Window().Call("open", "https://github.com/mlctrez/goapp-mdc")
	})

	topBar.Actions = []app.HTMLButton{reloadButton, codeButton, githubButton}

	body := app.Div().Body(
		nav,
		app.Div().Class("mdc-drawer-app-content").Body(
			&AppUpdateBanner{},
			topBar,
			app.Div().Class("main-content").ID("main-content").Body(
				topBar.Main().Body(
					app.Div().Style("display", "flex").Body(pageContent...),
				),
			),
		),
	)

	return body
}

func FlexGrid(cells ...app.UI) app.UI {
	return layout.Grid().Body(
		layout.Inner().Style("display", "flex").Body(cells...),
	)
}
