package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/demo"
)

func main() {
	demo.Routes()
	app.RunWhenOnBrowser()
	httpServer()
}

