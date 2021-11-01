package older

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/button"
)

type Index struct {
	app.Compo
}

func (i *Index) Render() app.UI {
	return app.Div().Class("goapp-mdc-coming-soon").Body(
		app.Div().Class("coming-soon-buttons").Body(
			&button.Button{Id: "btn1", Label: "Coming", Outlined: true},
			&button.Button{Id: "btn2", Label: "Soon", Raised: true, Callback: func(button app.HTMLButton) {
				button.OnClick(func(ctx app.Context, e app.Event) {
					ctx.Navigate("/demo")
				})
			}},
		),
		&PulseImage{},
	)
}

func (i *Index) OnPreRender(ctx app.Context) {
	i.initPage(ctx)
}

func (i *Index) OnNav(ctx app.Context) {
	i.initPage(ctx)
}

func (i *Index) initPage(ctx app.Context) {
	page := ctx.Page()
	page.SetAuthor("mlctrez")
	page.SetKeywords("go, golang, go-app, pwa, wasm, material design components")
	page.SetTitle("goapp-mdc")
	page.SetDescription("Material Design Components for go-app")
}
