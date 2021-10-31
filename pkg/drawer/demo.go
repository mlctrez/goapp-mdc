package drawer

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

type Demo struct {
	app.Compo
	base.JsUtil
}

func (d *Demo) Render() app.UI {

	navItems := list.Items{
		&list.Item{Text: "Inbox", Graphic: icon.MIInbox},
		&list.Item{Text: "Outgoing", Graphic: icon.MISend},
		//&list.Item{Type: list.ItemTypeDivider},
		&list.Item{Text: "Drafts", Graphic: icon.MIDrafts},
		&list.Item{Text: "Settings", Graphic: icon.MISettings},
		&list.Item{Text: "Ramen", Graphic: icon.MIRamenDining},
	}
	navItems.Select(0)

	return &Drawer{Id: d.UUID(), Type: Standard, List: &list.List{Id: "navigation", Type: list.Navigation, Items: navItems.UIList()}}
}

func (d *Demo) OnMount(ctx app.Context) {
	ctx.Handle(string(list.Select), d.eventHandler)
}

func (d *Demo) eventHandler(ctx app.Context, action app.Action) {
	log.Println("you clicked on item", action.Tags.Get("index"))
}
