package icon

import "github.com/maxence-charriere/go-app/v9/pkg/app"

const MaterialIconsClass = "material-icons"

func (m MaterialIcon) Span() app.HTMLSpan {
	return app.Span().Class(MaterialIconsClass).Text(m)
}

func (m MaterialIcon) I() app.HTMLI {
	return app.I().Class(MaterialIconsClass).Text(m)
}
