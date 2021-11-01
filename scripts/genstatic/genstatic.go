package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/demo"
)

func main() {
	demo.Routes()
	err := app.GenerateStaticWebsite("static", demo.BuildHandler())
	if err != nil {
		log.Fatalln(err)
	}
}
