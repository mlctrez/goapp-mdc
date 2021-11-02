//go:build !wasm

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/demo"
)

// GitCommit is set using : go build -ldflags "-X main.GitCommit=$(GIT_COMMIT)"
var GitCommit string

func httpServer() {
	handler := setupVersion(demo.BuildHandler())
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Println(err)
	}
}

func setupVersion(handler *app.Handler) *app.Handler {
	flag.Parse()
	switch flag.Arg(0) {
	case "dev":
		fmt.Println("using dynamic version")
	default:
		handler.Version = GitCommit
	}
	return handler
}
