package progress

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func ExampleNewLinear() {
	bar := NewLinear("linearProgress")
	bar.WithOnMount(func(ctx app.Context) {
		ctx.Async(func() {
			bar.Open()
			//bar.Determinate(false)
			var p float64
			for p = 0; p < 1; p += 0.01 {
				bar.SetProgress(p)
				time.Sleep(40 * time.Millisecond)
			}
			bar.SetProgress(1)
			time.Sleep(1500 * time.Millisecond)
			bar.Close()
		})
	})
}
