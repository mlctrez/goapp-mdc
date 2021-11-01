package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/demo/older"
)

func Routes() {
	app.Route("/", &Index{})
	app.Route("/ramen", &Ramen{})
	app.Route("/demo", &older.Demo{})
}

