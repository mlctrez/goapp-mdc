package demo

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/demo/markup"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/drawer"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

type CodeDemo struct {
	app.Compo
	base.JsUtil
	Content string
	Active  int
	list *list.List
}

func sortedNames() []string {
	var sortedNames []string
	for n := range markup.Code {
		sortedNames = append(sortedNames, n)
	}
	// reverse sort here
	sort.Slice(sortedNames, func(i, j int) bool { return sortedNames[i] > sortedNames[j] })
	return sortedNames
}

func (d *CodeDemo) OnNav(ctx app.Context) {
	d.LogWithP(d, "OnNav")
	url := ctx.Page().URL()
	idx, err := strconv.Atoi(url.Fragment)
	if err != nil {
		idx = 0
	}
	d.Active = idx
	d.Content = markup.Code[sortedNames()[d.Active]]
	d.list.Select(d.Active)
	d.Update()
	ctx.Defer(func(context app.Context) {
		app.Window().Get("Prism").Call("highlightAll")
	})
}

func (d *CodeDemo) Render() app.UI {
	d.LogWithPf(d, "Render active=%d", d.Active)
	if d.list == nil {
		d.LogWithP(d, "new list")
		items := list.Items{}
		for i, name := range sortedNames() {
			items = append(items, &list.Item{Text: name,
				Type: list.ItemTypeAnchor, Href: fmt.Sprintf("/code#%d", i)})
		}
		d.list = &list.List{Id: "codeNav", Type: list.Navigation, Items: items.UIList()}
	}

	body := &drawer.Drawer{Id: "codeNavigation", Type: drawer.Standard, List: d.list}

	if d.Content == "" {
		d.Content = markup.Code[sortedNames()[0]]
	}

	return PageBody(body, app.Raw(d.Content))
}

func (d *CodeDemo) OnMount(ctx app.Context) {
	d.LogWithPf(d, "OnMount active=%d", d.Active)
	ctx.Handle(string(list.Select), d.eventHandler)
	app.Window().Get("Prism").Call("highlightAll")
}

func (d *CodeDemo) eventHandler(ctx app.Context, action app.Action) {
	if selectedIndex, err := strconv.Atoi(action.Tags.Get("index")); err != nil {
		return
	} else {
		ctx.Navigate(fmt.Sprintf("/code#%d", selectedIndex))
	}
}
