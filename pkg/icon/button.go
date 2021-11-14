package icon

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
)

// Button represents a material icon-button without toggle.
type Button struct {
	app.Compo
	autoinit.AutoInit
	Icon      MaterialIcon
	Id        string
	AriaLabel string
	api       app.Value
}

func (b *Button) Render() app.UI {
	button := app.Button().Class("mdc-icon-button")
	autoinit.MDCRipple.DataMdcAutoInitButton(button)
	if b.Id != "" {
		button.ID(b.Id)
	}
	button.Class("material-icons")
	button.Aria("label", b.AriaLabel)
	button.Body(RippleDiv(), app.Text(b.Icon))
	return button
}

func (b *Button) OnMount(ctx app.Context) {
	ctx.Defer(func(context app.Context) {
		b.api = b.AutoInitComponent(b.JSValue(), autoinit.MDCRipple)
		b.api.Set("unbounded", true)
	})
}
