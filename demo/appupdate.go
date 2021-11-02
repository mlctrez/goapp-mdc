package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/banner"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
)

// AppUpdateBanner demonstrates how to wrap banner.Banner to handle go-app OnAppUpdate.
type AppUpdateBanner struct {
	app.Compo
	base.JsUtil
	bnr *banner.Banner
}

func (d *AppUpdateBanner) Render() app.UI {
	if d.bnr == nil {
		d.bnr = &banner.Banner{
			Id: "appUpdateBanner", Fixed: true, Centered: true,
			Text: "A new version is available, would you like to install?",
		}
		d.bnr.Buttons = d.bannerButtons()
	}
	return d.bnr
}

func (d *AppUpdateBanner) bannerButtons() []app.UI {
	primary := &button.Button{Id: "updateBannerYes", Label: "yes",
		Icon: string(icon.MIUpdate), Banner: true, BannerAction: "primary"}
	secondary := &button.Button{Id: "updateBannerNo", Label: "later",
		Icon: string(icon.MIWatchLater), Banner: true, BannerAction: "secondary"}
	return []app.UI{primary, secondary}
}

func (d *AppUpdateBanner) onBannerClose(ctx app.Context, reason string) {
	switch reason {
	case "primary": // Yes button
		ctx.Reload()
	case "secondary": // Later button
		// This could SetState for a future time to ask
	}
}

func (d *AppUpdateBanner) OnMount(ctx app.Context) {
	d.bnr.ActionClose(ctx, d.onBannerClose)
}

func (d *AppUpdateBanner) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		d.bnr.ActionOpen(ctx)
	}
}
