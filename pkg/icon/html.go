package icon

import "github.com/maxence-charriere/go-app/v9/pkg/app"

const MaterialIconsClass = "material-icons"
const MaterialIconButton = "mdc-icon-button"

func (m MaterialIcon) Span() app.HTMLSpan {
	return m.SpanClasses(MaterialIconsClass)
}

func (m MaterialIcon) SpanClasses(classes ...string) app.HTMLSpan {
	return app.Span().Class(classes...).Text(m)
}

func (m MaterialIcon) I() app.HTMLI {
	return m.IClasses(MaterialIconsClass)
}

func (m MaterialIcon) IClasses(classes ...string) app.HTMLI {
	return app.I().Class(classes...).Text(m)
}

func (m MaterialIcon) Button() app.HTMLButton {
	return m.ButtonClasses(MaterialIconsClass, MaterialIconButton)
}

func (m MaterialIcon) ButtonClasses(classes ...string) app.HTMLButton {
	return app.Button().Class(classes...).Text(m)
}

func (m MaterialIcon) IItemGraphic() app.HTMLI {
	return m.IClasses(MaterialIconsClass, "mdc-deprecated-list-item__graphic").Aria("hidden", "true")
}
