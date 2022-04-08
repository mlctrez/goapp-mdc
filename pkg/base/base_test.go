package base

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func TestHTMLButtonUpdate(t *testing.T) {
	// should not fail
	HTMLButtonUpdate([]app.HTMLButton{nil}, nil)
}
