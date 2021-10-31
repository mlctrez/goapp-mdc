package icon

import "github.com/maxence-charriere/go-app/v9/pkg/app"

const MaterialIconsClass = "material-icons"

func (m MaterialIcon) Span() app.HTMLSpan {
	return app.Span().Class(MaterialIconsClass).Text(m)
}

func (m MaterialIcon) I() app.HTMLI {
	return m.IClasses(MaterialIconsClass)
}

func (m MaterialIcon) IClasses(classes ...string) app.HTMLI {
	return app.I().Class(classes...).Text(m)
}

func (m MaterialIcon) IItemGraphic() app.HTMLI {
	return m.IClasses(MaterialIconsClass, "mdc-deprecated-list-item__graphic").Aria("hidden", "true")
}
