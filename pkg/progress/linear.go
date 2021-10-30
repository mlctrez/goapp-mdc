package progress

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

type Linear struct {
	app.Compo
	base.JsUtil
	id       string
	label    string
	progress float64
	onmount  func(ctx app.Context)
	target   app.Value
}

func NewLinear(id string) *Linear {
	return &Linear{id: id, target: app.Undefined()}
}

func (b *Linear) Label(label string) *Linear {
	b.label = label
	return b
}

func (b *Linear) WithOnMount(cb func(ctx app.Context)) *Linear {
	b.onmount = cb
	return b
}

func (b *Linear) Render() app.UI {

	root := app.Div().ID(b.id).Class("mdc-linear-progress", "mdc-linear-progress--closed").Aria("hidden", "true")
	root.Aria("valuemin", "0").Aria("valuemax", "1").Aria("valuenow", "0")
	if b.label != "" {
		root.Aria("label", b.label)
	}
	root.Body(
		app.Div().Class("mdc-linear-progress__buffer").Body(
			app.Div().Class("mdc-linear-progress__buffer-bar"),
			app.Div().Class("mdc-linear-progress__buffer-dots"),
		),
		app.Div().Class("mdc-linear-progress__bar", "mdc-linear-progress__primary-bar").Body(
			app.Span().Class("mdc-linear-progress__bar-inner"),
		),
		app.Div().Class("mdc-linear-progress__bar", "mdc-linear-progress__secondary-bar").Body(
			app.Span().Class("mdc-linear-progress__bar-inner"),
		),
	)

	return root
}

func (b *Linear) OnMount(ctx app.Context) {

	b.target = b.JsNewAtPath("mdc.linearProgress.MDCLinearProgress", app.Window().GetElementByID(b.id))

	if b.onmount != nil {
		b.onmount(ctx)
	}
}

func (b *Linear) Open() {
	if !b.target.IsUndefined() {
		b.target.Call("open")
	}
}

func (b *Linear) Determinate(d bool) {
	if !b.target.IsUndefined() {
		b.target.Set("determinate", d)
	}
}

func (b *Linear) SetProgress(f float64) {
	b.progress = f
	if !b.target.IsUndefined() {
		b.target.Set("progress", b.progress)
	}
}
func (b *Linear) Close() {
	if !b.target.IsUndefined() {
		b.target.Call("close")
	}
}

/*


	bar := progress.NewLinear("progress")
	bar.WithOnMount(func(ctx app.Context) {
		ctx.Async(func() {
			bar.Open()
			bar.Determinate(false)
			var p float64
			for p = 0; p < 1; p += 0.01 {
				//bar.SetProgress(p)
				time.Sleep(40 * time.Millisecond)
			}
			bar.Close()
		})
	})



*/
