package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/drawer"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

var NavigationItems []*list.Item

type Navigation struct {
	app.Compo
	base.JsUtil
	Type   drawer.Type
	drawer *drawer.Drawer
	items  list.Items
	//list  *list.List
}

func (n *Navigation) Render() app.UI {
	if n.items == nil {
		n.items = make(list.Items, len(NavigationItems))
		for i, item := range NavigationItems {
			n.items[i] = &list.Item{
				Type:      item.Type,
				Graphic:   item.Graphic,
				Text:      item.Text,
				Secondary: item.Secondary,
				Href:      item.Href,
			}
		}
		n.drawer = &drawer.Drawer{
			Type: n.Type, Id: "navigationDrawer",
			List: &list.List{Type: list.Navigation,
				Id: "navigationList", Items: n.items.UIList(),
			},
		}
	}
	return n.drawer
}

func (n *Navigation) OnMount(ctx app.Context) {
	n.items.SelectHref(ctx.Page().URL().Path)
}
