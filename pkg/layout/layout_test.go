package layout

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func ExampleGrid() {

	Grid().Body(Inner().Body(
		Cell().Body(app.Input()),
		Cell().Body(app.Br()),
		CellModified("top", 4).Body(app.Textarea()),
	))

}
