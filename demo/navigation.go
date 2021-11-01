package demo

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/drawer"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

type Navigation struct {
	app.Compo
	base.JsUtil
	items list.Items
	list  *list.List
}

func (n *Navigation) Render() app.UI {
	return &drawer.Drawer{Type: drawer.Standard, Id: "navigationDrawer", List: n.list}
}

func (n *Navigation) OnMount(ctx app.Context) {
	if n.items == nil {
		n.items = list.Items{
			&list.Item{Text: "Home", Graphic: icon.MIHome, Type: list.ItemTypeAnchor, Href: "/"},
			&list.Item{Text: "Ramen", Graphic: icon.MIRamenDining, Type: list.ItemTypeAnchor, Href: "/ramen"},
		}
		n.items.SelectHref(ctx.Page().URL().Path)

		n.list = &list.List{Type: list.Navigation, Id: "navigationList", Items: n.items.UIList()}
	}
	ctx.Handle(string(list.Select), n.eventHandler)
}

func (n *Navigation) eventHandler(ctx app.Context, action app.Action) {
	if action.Value != n.list {
		return
	}
	log.Println("you clicked on item", action.Tags.Get("index"))
}
