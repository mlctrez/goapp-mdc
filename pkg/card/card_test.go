package card

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

func TestCard_Render(t *testing.T) {
	id := "testID"
	compo := &Card{Id: id}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Div().ID(compo.Id).Class("mdc-card"))
}
