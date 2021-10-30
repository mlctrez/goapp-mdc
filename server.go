//go:build !wasm

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var GitCommit string

type resourceProvider struct {
}

func (r resourceProvider) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if strings.HasPrefix(request.RequestURI, "/web/") {
		file, err := os.Open(request.RequestURI[1:])
		if os.IsNotExist(err) {
			writer.WriteHeader(404)
			return
		}
		ext := filepath.Ext(request.RequestURI)

		contentType := mime.TypeByExtension(ext)
		writer.Header().Set("Content-Type", contentType)
		if ext == ".wasm" {
			writer.Header().Set("Content-Encoding", "br")
		}

		io.Copy(writer, file)
		return
	}

	fmt.Println("404")
	writer.WriteHeader(404)
}

func (r resourceProvider) Package() string { return "" }
func (r resourceProvider) Static() string  { return "" }
func (r resourceProvider) AppWASM() string { return "/web/app.wasm" }

var _ app.ResourceProvider = (*resourceProvider)(nil)

func buildHandler() (h *app.Handler, err error) {
	var open io.ReadCloser
	if open, err = os.Open("handler.json"); err != nil {
		return
	}
	defer open.Close()
	h = &app.Handler{}
	err = json.NewDecoder(open).Decode(h)
	h.Resources = &resourceProvider{}

	h = &app.Handler{
		Author:          "mlctrez",
		Description:     "Material Design Components for go-app",
		Icon:            app.Icon{Default: "/web/logo-192.png", Large: "/web/logo-512.png"},
		Resources:       &resourceProvider{},
		Name:         "MDC for go-app",
		Scripts: []string{
			//"https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js",
			"/web/material-components-web.min.js",
			//"/web/app.js",
			"/web/prism.js",
		},
		ShortName: "goapp-mdc",
		Styles: []string{
			//"https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css",
			"/web/material-components-web.min.css",
			"/web/style.css",
			"/web/prism.css",
			"https://fonts.googleapis.com/icon?family=Material+Icons"},
		//ThemeColor: "#4d60b8",
		Title:      "Material Design Components for go-app",
	}

	return
}

func httpServer() {

	//`<script src="https://www.google.com/recaptcha/enterprise.js?render=6Ldt8sgcAAAAACwJjJMaRH3b31xDXBB6IYvBpLmc"></script>`,
	//`<style>.grecaptcha-badge { visibility: hidden; }</style>`,

	handler, err := buildHandler()
	if err != nil {
		panic(err)
	}

	if os.Getenv("GEN_STATIC") != "" {
		err = app.GenerateStaticWebsite("static", handler)
		if err != nil {
			panic(err)
		}
		return
	}

	if os.Getenv("VERSION") == "dynamic" {
		log.Println("using dynamic version")
		handler.Version = ""
	} else {
		log.Println("VERSION=", handler.Version)
	}

	if err = http.ListenAndServe(":8000", handler); err != nil {
		log.Println(err)
	}

}
