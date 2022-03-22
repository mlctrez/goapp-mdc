package progress

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func ExampleNewCircular() {
	circular := NewCircular("largeCircularProgress", 48)
	circular.WithOnMount(func(ctx app.Context) {
		ctx.Async(func() {
			circular.Determinate(true)
			circular.Open()
			for {
				var p float64
				for p = 0; p < 1; p += 0.01 {
					circular.SetProgress(p)
					time.Sleep(40 * time.Millisecond)
				}
				circular.SetProgress(1)
			}
		})
	})
}
