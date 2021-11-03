package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/drawer"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

// TODO: make this NavigationItems immutable

var NavigationItems list.Items

type Navigation struct {
	app.Compo
	base.JsUtil
	items list.Items
	list  *list.List
}

func (n *Navigation) Render() app.UI {
	if n.items == nil {
		n.items = NavigationItems
		n.list = &list.List{Type: list.Navigation, Id: "navigationList", Items: n.items.UIList()}
	}
	return &drawer.Drawer{Type: drawer.Standard, Id: "navigationDrawer", List: n.list}
}

func (n *Navigation) OnMount(ctx app.Context) {
	n.items.SelectHref(ctx.Page().URL().Path)
}
