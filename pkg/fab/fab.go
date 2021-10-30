package fab

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
)

type Fab struct {
	app.Compo
	base.JsUtil
	Id       string
	Icon     icon.MaterialIcon
	Label    string
	Extended bool
	Mini     bool
	callback func(button app.HTMLButton)
	onmount  func(ctx app.Context)
}

func New(id string, icon icon.MaterialIcon) *Fab {
	return &Fab{Id: id, Icon: icon}
}

func (b *Fab) WithCallback(cb func(button app.HTMLButton)) *Fab {
	b.callback = cb
	return b
}

func (b *Fab) WithOnMount(cb func(ctx app.Context)) *Fab {
	b.onmount = cb
	return b
}

func (b *Fab) OnMount(ctx app.Context) {
	value := b.JsValueAtPath("mdc.ripple.MDCRipple")
	if !value.IsUndefined() {
		value.Call("attachTo", app.Window().GetElementByID(b.Id))
	}
	if b.onmount != nil {
		b.onmount(ctx)
	}
}

func (b *Fab) Render() app.UI {
	result := app.Button().ID(b.Id)

	icon := app.Span().Class("mdc-fab__icon", "material-icons").Text(b.Icon)
	ripple := app.Div().Class("mdc-fab__ripple")

	if !b.Extended && !b.Mini {
		// regular fab
		result.Class("mdc-fab").Aria("label", b.Label).Body(ripple, icon)
	} else {
		if b.Mini {
			result.Class("mdc-fab", "mdc-fab--mini").Body(ripple, icon)
		} else {
			result.Class("mdc-fab", "mdc-fab--extended").Body(
				ripple,
				icon,
				app.Span().Class("mdc-fab__label").Text(b.Label),
			)
		}
	}

	if b.callback != nil {
		b.callback(result)
	}

	return result
}
