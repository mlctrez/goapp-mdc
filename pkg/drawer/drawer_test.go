package drawer

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

func TestDrawer_Render(t *testing.T) {
	id := "testID"
	compo := &Drawer{Id: id}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

	ct.At(0).Match(app.Aside().ID(compo.Id).Class("mdc-drawer"))
	ct.At(0, 0).Match(app.Div().Class("mdc-drawer__content"))




}
