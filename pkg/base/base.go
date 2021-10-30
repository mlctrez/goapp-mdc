package base

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type JsUtil struct{}

func (b *JsUtil) JsValueAtPath(path string) app.Value {
	var current app.Value = app.Window()
	for _, part := range strings.Split(path, ".") {
		current = current.Get(part)
		if current.IsUndefined() {
			fmt.Printf("jsutil.getValue(%q) : undefined at %q\n", path, part)
			break
		}
	}
	return current
}

func (b *JsUtil) JsNewAtPath(path string, args ...interface{}) app.Value {
	v := b.JsValueAtPath(path)
	if v.Truthy() {
		return v.New(args...)
	}
	return app.Undefined()
}

// MDCRipple constructs a new mdc.ripple.MDCRipple with the element
func (b *JsUtil) MDCRipple(id string) app.Value {
	elemId := app.Window().GetElementByID(id)
	return b.JsNewAtPath("mdc.ripple.MDCRipple", elemId)
}

// UUID generates a new uuid string for testing and demo purposes
func (b *JsUtil) UUID() string {
	return uuid.New().String()
}

type Counter struct {
	app.Compo
	Label string
	Count int
}

func (c *Counter) Render() app.UI {
	return app.Span().Text(fmt.Sprintf("counter %q is currently %d", c.Label, c.Count))
}
