package demo

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/drawer"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

type DrawerDemo struct {
	app.Compo
	base.JsUtil
}

func (d *DrawerDemo) Render() app.UI {

	navItems := list.Items{
		&list.Item{Text: "Inbox", Graphic: icon.MIInbox},
		&list.Item{Text: "Outgoing", Graphic: icon.MISend},
		&list.Item{Type: list.ItemTypeDivider},
		&list.Item{Text: "Drafts", Graphic: icon.MIDrafts},
		&list.Item{Text: "Settings", Graphic: icon.MISettings},
		&list.Item{Text: "Ramen", Graphic: icon.MIRamenDining},
	}
	navItems.Select(0)

	body := &drawer.Drawer{Id: d.UUID(), Type: drawer.Standard, List: &list.List{Id: "navigation", Type: list.Navigation, Items: navItems.UIList()}}
	return PageBody(body)
}

func (d *DrawerDemo) OnMount(ctx app.Context) {
	ctx.Handle(string(list.Select), d.eventHandler)
}

func (d *DrawerDemo) eventHandler(ctx app.Context, action app.Action) {
	log.Println("you clicked on item", action.Tags.Get("index"))
}
