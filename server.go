//go:build !wasm

package main

import (
	"log"
	"net/http"

	"github.com/mlctrez/goapp-mdc/demo"
)

var GitCommit string

func httpServer() {
	handler := demo.BuildHandler()
	handler.Version = GitCommit
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Println(err)
	}
}
