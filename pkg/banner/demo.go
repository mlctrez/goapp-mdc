package banner

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
)

type Demo struct {
	app.Compo
	banner   *Banner
	open     *button.Button
	close    *button.Button
	centered *checkbox.Checkbox
	fixed    *checkbox.Checkbox
}

func buildButton(action string, callback func(ctx app.Context, e app.Event)) *button.Button {
	return &button.Button{
		Id: action, Label: action, Raised: true, Callback: func(button app.HTMLButton) {
			button.OnClick(callback)
		},
	}
}

func buildBannerButton(action string, callback func(ctx app.Context, e app.Event)) *button.Button {
	return &button.Button{
		Id: "bannerButton-" + action, Label: action, Banner: true, BannerAction: action, Callback: func(button app.HTMLButton) {
			button.OnClick(callback)
		},
	}
}

func buildCheckbox(action string, onChange func(bool)) *checkbox.Checkbox {
	return &checkbox.Checkbox{Id: action, Label: action, Callback: func(input app.HTMLInput) {
		input.OnChange(func(ctx app.Context, e app.Event) {
			onChange(ctx.JSSrc().Get("checked").Bool())
		})
	}}
}

func (c *Demo) Render() app.UI {

	if c.banner == nil {
		c.banner = &Banner{Id: "demoBanner", Text: "banner text goes here",
			OnClosed: func(s string) {
				fmt.Println("banner was closed with reason", s)
			},
			Buttons: []app.UI{
				buildBannerButton("primary", func(ctx app.Context, e app.Event) {
					fmt.Println("you clicked the primary action")
				}),
				buildBannerButton("secondary", func(ctx app.Context, e app.Event) {
					fmt.Println("you clicked the secondary action")
				}),
			}}
		c.open = buildButton("open", func(ctx app.Context, e app.Event) { c.banner.Open() })
		c.close = buildButton("close", func(ctx app.Context, e app.Event) { c.banner.Close() })

		c.centered = buildCheckbox("centered", func(b bool) {
			fmt.Println("setting centered to", b)
			c.banner.Centered = b
			c.banner.Update()
		})
		c.fixed = buildCheckbox("fixed", func(b bool) {
			fmt.Println("setting fixed to", b)
			c.banner.Fixed = b
			c.banner.Update()
		})

	}

	return app.Div().Body(
		c.banner,
		app.Hr(),
		app.Div().Text("banner controls"),
		app.Br(),
		c.open, c.close, c.centered, c.fixed)

}
