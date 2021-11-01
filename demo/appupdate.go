package demo

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/banner"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
)

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
	log.Println("banner was closed with reason", reason)
	switch reason {
	case "primary":
		ctx.Reload()
	case "secondary":
		// set a timer to open again in X hours?
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
