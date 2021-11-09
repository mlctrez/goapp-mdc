package main

import (
	"bytes"
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/bar"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
)

func main() {
	b := &bar.TopAppBar{}
	b.Title = "Title goes here"
	b.Navigation = []app.HTMLButton{icon.MIMenu.Button()}
	b.Actions = []app.HTMLButton{icon.MIFavorite.Button()}
	html := &bytes.Buffer{}
	app.PrintHTMLWithIndent(html, b)
	fmt.Println(html.String())
}
