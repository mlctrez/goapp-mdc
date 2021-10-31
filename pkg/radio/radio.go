package radio

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Radio struct {
	app.Compo
	Id string
}

func (r *Radio) Render() app.UI {
	return app.Div()
}
func (r *Radio) OnMount(_ app.Context) {
}
