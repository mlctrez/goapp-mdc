package progress

import (
	"fmt"
	"testing"
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

func TestCircular_DeterminateSVG(t *testing.T) {
	c := Circular{size: 48}
	svg := c.DeterminateSVG()
	fmt.Println(svg)
}


var example = `<svg class="mdc-circular-progress__determinate-circle-graphic" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
<circle class="mdc-circular-progress__determinate-track" cx="24" cy="24" r="18.0" stroke-width="4.0"/>
<circle class="mdc-circular-progress__determinate-circle" cx="24" cy="24" r="18.0" stroke-dasharray="113.1" stroke-dashoffset="113.1" stroke-width="4.0"/>
</svg>`
