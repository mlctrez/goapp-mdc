package radio

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

func TestRadio_Render(t *testing.T) {
	id := "testID"
	compo := &Radio{Id: id}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

}
