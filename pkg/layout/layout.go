/*
Package layout implements the markup defined in

https://github.com/material-components/material-components-web/tree/master/packages/mdc-layout-grid

See ExampleGrid for usage.

*/
package layout

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Grid is the mandatory layout grid element with class mdc-layout-grid.
func Grid() app.HTMLDiv {
	return app.Div().Class("mdc-layout-grid")
}

// Inner is the mandatory wrapping grid cell with class mdc-layout-grid__inner.
func Inner() app.HTMLDiv {
	return app.Div().Class("mdc-layout-grid__inner")
}

// Cell is the mandatory layout grid cell with class mdc-layout-grid__cell.
// Use CellModified to specify alignment and span width.
func Cell() app.HTMLDiv {
	return app.Div().Class("mdc-layout-grid__cell")
}

// CellModified is the mandatory innermost layout div with class mdc-layout-grid__cell.
// It allows optional classes for alignment ( top, middle, bottom ) and span width ( 1-12 ).
// If align and span values are out of range, no additional classes are added.
func CellModified(align string, span int) app.HTMLDiv {
	classes := []string{"mdc-layout-grid__cell"}
	switch align {
	case "top", "middle", "bottom":
		classes = append(classes, "mdc-layout-grid__cell--align-"+align)
	}
	if span > 0 && span < 13 {
		classes = append(classes, fmt.Sprintf("mdc-layout-grid__cell--span-%d", span))
	}
	return app.Div().Class(classes...)
}

// CellWide is equivalent to CellModified("middle", 12)
func CellWide() app.HTMLDiv {
	return CellModified("middle", 12)
}
