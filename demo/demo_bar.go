package demo

import (
	"fmt"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/bar"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
)

type BarDemo struct {
	app.Compo
}

func (d *BarDemo) Render() app.UI {

	topBar := &bar.TopAppBar{Title: "Top App Bar Demo"}
	topBar.Navigation = []app.HTMLButton{icon.MIMenu.Button().OnClick(base.OnEvent("navigation click"))}

	codeButton := icon.MICode.Button().OnClick(func(ctx app.Context, e app.Event) {
		fragment := strings.TrimPrefix(ctx.Page().URL().Path, "/") + ".go"
		ctx.Navigate(fmt.Sprintf("/code#%s", fragment))
	})

	githubButton := app.Button().Class(icon.MaterialIconsClass, icon.MaterialIconButton).
		Body(app.Raw(GitHubSvg)).OnClick(func(ctx app.Context, e app.Event) {
		app.Window().Call("open", "https://github.com/mlctrez/goapp-mdc")
	})

	topBar.Actions = []app.HTMLButton{codeButton, githubButton}

	return app.Div().Style("display", "flex").Body(
		&AppUpdateBanner{}, &Navigation{},
		)

	//return app.Div().Body(
	//	topBar,
	//	app.Main().Class("mdc-top-app-bar--fixed-adjust").Body(
	//		&AppUpdateBanner{},
	//		&Navigation{},
	//	),
	//)
}

const GitHubSvg = `<svg class="mdc-button__icon" width="32" height="32" aria-hidden="true" viewBox="0 0 16 16">
    <path fill-rule="evenodd" d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"></path>
</svg>`
