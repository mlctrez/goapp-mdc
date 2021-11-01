package demo

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func BuildHandler() *app.Handler {
	return &app.Handler{
		Author:          "mlctrez",
		Description:     "Material Design Components for go-app",
		Icon:            app.Icon{Default: "/web/logo-192.png", Large: "/web/logo-512.png"},
		Name:            "MDC for go-app",
		BackgroundColor: "#111",
		Scripts: []string{
			//"https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js",
			"/web/material-components-web.min.js",
		},
		ShortName: "goapp-mdc",
		Styles: []string{
			"https://fonts.googleapis.com/icon?family=Material+Icons",
			//"https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css",
			"/web/material-components-web.min.css",
			"/web/style.css",
		},
		Title: "Material Design Components for go-app",
	}
}
