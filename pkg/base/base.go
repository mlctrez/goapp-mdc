package base

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type JsUtil struct{}

func (b *JsUtil) JsValueAtPath(path string) app.Value {
	return b.JsValueAt(app.Window(), path, true)
}

func (b *JsUtil) JsValueAt(root app.Value, path string, warn bool) app.Value {
	if root == nil || root.IsUndefined() {
		return app.Undefined()
	}
	var current = root
	for _, part := range strings.Split(path, ".") {
		current = current.Get(part)
		if current.IsUndefined() {
			if warn {
				fmt.Printf("jsutil.getValue(%q) : undefined at %q\n", path, part)
			}
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

// MDCRipple constructs a new mdc.ripple.MDCRipple from a provided element id
func (b *JsUtil) MDCRipple(id string) app.Value {
	return b.MDCRippleVal(app.Window().GetElementByID(id))
}

// MDCRippleVal constructs a new mdc.ripple.MDCRipple from a provided app.Value
func (b *JsUtil) MDCRippleVal(val app.Value) app.Value {
	if val == nil || val.IsUndefined() {
		return app.Undefined()
	}
	return b.JsNewAtPath("mdc.ripple.MDCRipple", val)
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
