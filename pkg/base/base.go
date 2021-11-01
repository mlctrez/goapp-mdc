package base

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type JsUtil struct{}

func (j *JsUtil) JsValueAtPath(path string) app.Value {
	return j.JsValueAt(app.Window(), path, true)
}

func (j *JsUtil) JsValueAt(root app.Value, path string, warn bool) app.Value {
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

func (j *JsUtil) JsNewAtPath(path string, args ...interface{}) app.Value {
	v := j.JsValueAtPath(path)
	if v.Truthy() {
		return v.New(args...)
	}
	return app.Undefined()
}

// MDCRipple constructs a new mdc.ripple.MDCRipple from a provided element id
func (j *JsUtil) MDCRipple(id string) app.Value {
	return j.MDCRippleVal(app.Window().GetElementByID(id))
}

// MDCRippleVal constructs a new mdc.ripple.MDCRipple from a provided app.Value
func (j *JsUtil) MDCRippleVal(val app.Value) app.Value {
	if val == nil || val.IsUndefined() {
		return app.Undefined()
	}
	return j.JsNewAtPath("mdc.ripple.MDCRipple", val)
}

// UUID generates a new uuid string for testing and demo purposes
func (j *JsUtil) UUID() string {
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

func (j *JsUtil) LogWithP(o interface{}, msg string) {
	ptr := reflect.ValueOf(o).Pointer()
	typ := reflect.TypeOf(o)
	log.Printf("%s(%d) %s\n", typ, ptr, msg)
}

func (j *JsUtil) LogWithPf(o interface{}, msg string, args ...interface{}) {
	j.LogWithP(o, fmt.Sprintf(msg, args))
}
