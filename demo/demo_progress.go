package demo

import (
	"fmt"
	"reflect"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/progress"
)

type ProgressDemo struct {
	app.Compo
	base.JsUtil
	circular       *progress.Circular
	circularInd    *progress.Circular
	circularColors *progress.Circular
	linear         *progress.Linear
	linearInd      *progress.Linear
}

func (d *ProgressDemo) Render() app.UI {

	if d.circular == nil {
		d.circular = progress.NewCircular("circularProgress", 48).Label("wait for it")
		d.circularInd = progress.NewCircular("circularProgressInd", 48).Label("wait for it ind")
		d.circularColors = progress.NewCircular("circularProgressColors", 48).Label("wait for it color")
		d.circularColors.Colors([4]string{"red", "green", "blue", "cyan"})
		d.linear = progress.NewLinear("linearProgress").Label("wait for it pt 2")
		d.linearInd = progress.NewLinear("linearProgressInd").Label("wait for it pt 2 ind")
	}

	body := layout.Grid().Body(
		row("Circular", d.circular, d.showButton(d.circular, true)),
		row("Circular Indeterminate", d.circularInd, d.showButton(d.circularInd, false)),
		row("Circular Colors", d.circularColors, d.showButton(d.circularColors, false)),
		row("Linear Progress", d.linear, d.showButton(d.linear, true)),
		row("Linear Indeterminate", d.linearInd, d.showButton(d.linearInd, false)),
	)

	return PageBody(body)
}

func row(text string, component app.UI, button app.UI) app.UI {
	return layout.Inner().Body(
		layout.CellModified("middle", 4).Body(app.Text(text)),
		layout.CellModified("bottom", 4).Style("height","50px").Body(component),
		layout.CellModified("middle", 4).Body(button),
	)
}

func (d *ProgressDemo) showButton(c progress.Api, determinate bool) app.UI {
	buttonId := fmt.Sprintf("button_%s_%t", reflect.TypeOf(c).Name(), determinate)
	return &button.Button{Id: buttonId, Label: "Show",
		Callback: func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				button.JSValue().Call("blur")
				go func() {
					c.Determinate(determinate)
					c.Open()
					for i := 0; i < 100; i++ {
						if determinate {
							c.SetProgress(float64(i) / float64(100))
						}
						time.Sleep(50 * time.Millisecond)
					}
					c.Close()
					time.Sleep(500 * time.Millisecond)
					if determinate {
						c.SetProgress(0)
					}
				}()
			})
		}}
}
