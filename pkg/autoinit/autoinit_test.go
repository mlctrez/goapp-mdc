package autoinit_test

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/autoinit"
)

type ExampleRippleComponent struct {
	app.Compo
	autoinit.AutoInit
}

func (e *ExampleRippleComponent) OnMount() {
	e.AutoInitComponent(e.JSValue(), autoinit.MDCRipple)
}

func ExampleAutoInit_AutoInitComponent() {
	// see OnMount() for example usage
}
