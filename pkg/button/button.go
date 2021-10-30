package button

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Button struct {
	app.Compo
	base.JsUtil
	Id           string
	Label        string
	Icon         string
	Raised       bool
	Unelevated   bool
	Outlined     bool
	TrailingIcon bool
	// button intended for use in dialog setting
	Dialog       bool
	DialogAction string
	// button intended for use in banner setting
	Banner       bool
	BannerAction string
	// button intended for use in card actions
	CardAction bool
	Callback   func(button app.HTMLButton)
}

func (b *Button) Render() app.UI {
	buttonLabel := app.Span().Class("mdc-button__label").Text(b.Label)

	buildIcon := func() app.UI {
		return app.I().Class("material-icons mdc-button__icon").Text(b.Icon)
	}()

	result := app.Button().ID(b.Id).Class(b.buildClasses()...).Body(
		app.Span().Class("mdc-button__ripple"),
		app.If(b.TrailingIcon,
			buttonLabel,
			app.If(b.Icon != "", buildIcon),
		).Else(
			app.If(b.Icon != "", buildIcon),
			buttonLabel,
		),
	)

	if b.Callback != nil {
		b.Callback(result)
	}

	if b.DialogAction != "" {
		result.DataSet("mdc-dialog-action", b.DialogAction)
	}

	return result
}

func (b *Button) OnMount(ctx app.Context) {
	value := b.JsValueAtPath("mdc.ripple.MDCRipple")
	if value.Truthy() {
		value.Call("attachTo", app.Window().GetElementByID(b.Id))
	}
}

func (b *Button) buildClasses() []string {
	var classes = []string{"mdc-button"}
	if b.Icon != "" {
		if b.TrailingIcon {
			classes = append(classes, "mdc-button--icon-trailing")
		} else {
			classes = append(classes, "mdc-button--icon-leading")
		}
	}
	if b.Dialog {
		classes = append(classes, "mdc-dialog__button")
	}
	if b.Banner {
		switch b.BannerAction {
		case "secondary", "primary":
			classes = append(classes, fmt.Sprintf("mdc-banner__%s-action", b.BannerAction))
		}
	}
	if b.CardAction {
		classes = append(classes, "mdc-card__action mdc-card__action--button")
	}

	if b.Raised {
		classes = append(classes, "mdc-button--raised")
	} else if b.Unelevated {
		classes = append(classes, "mdc-button--unelevated")
	} else {
		if b.Outlined {
			classes = append(classes, "mdc-button--outlined")
		}
	}

	return classes
}
