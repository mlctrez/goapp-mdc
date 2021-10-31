package list

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type GroupItem struct {
	SubHeader string
	List      *List
}

type Group struct {
	app.Compo
	Items []*GroupItem
}

func (g *Group) Render() app.UI {
	listGroup := app.Div()
	listGroup.Class("mdc-deprecated-list-group")
	var groupBody []app.UI
	for _, item := range g.Items {
		groupBody = append(groupBody, app.H3().Class("mdc-deprecated-list-group__subheader").Text(item.SubHeader))
		groupBody = append(groupBody, item.List)
	}
	listGroup.Body(groupBody...)
	return listGroup
}
