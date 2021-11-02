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
			"https://cdnjs.cloudflare.com/ajax/libs/material-components-web/13.0.0/material-components-web.min.js",
			"https://cdnjs.cloudflare.com/ajax/libs/prism/1.25.0/prism.min.js",
			"https://cdnjs.cloudflare.com/ajax/libs/prism/1.25.0/components/prism-go.min.js",
		},
		Env: map[string]string{
			"RECAPTCHA_SITE_KEY": "6Ldt8sgcAAAAACwJjJMaRH3b31xDXBB6IYvBpLmc",
		},
		ShortName: "goapp-mdc",
		Styles: []string{
			"https://fonts.googleapis.com/icon?family=Material+Icons",
			"https://fonts.googleapis.com/css2?family=Roboto&display=swap",
			"https://cdnjs.cloudflare.com/ajax/libs/material-components-web/13.0.0/material-components-web.min.css",
			"https://cdnjs.cloudflare.com/ajax/libs/prism-themes/1.9.0/prism-material-light.min.css",
			"/web/style.css",
		},
		Title: "Material Design Components for go-app",
	}
}
