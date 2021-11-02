package demo

import (
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
}

func sortedNames() []string {
	var sortedNames []string
	for n := range markup.Code {
		sortedNames = append(sortedNames, n)
	}
	sort.Slice(sortedNames, func(i, j int) bool { return sortedNames[i] > sortedNames[j] })
	return sortedNames
}

func (d *CodeDemo) Render() app.UI {

	navItems := list.Items{}

	for _, name := range sortedNames() {
		navItems = append(navItems, &list.Item{Text: name})
	}

	navItems.Select(0)

	body := &drawer.Drawer{Id: d.UUID(), Type: drawer.Standard, List: &list.List{Id: "codeNav", Type: list.Navigation, Items: navItems.UIList()}}

	if d.Content =="" {
		d.Content = markup.Code[sortedNames()[0]]
	}


	return PageBody(body, app.Raw(d.Content))
}

func (d *CodeDemo) OnMount(ctx app.Context) {
	ctx.Handle(string(list.Select), d.eventHandler)
	app.Window().Get("Prism").Call("highlightAll")
}

func (d *CodeDemo) eventHandler(ctx app.Context, action app.Action) {
	if selectedIndex, err := strconv.Atoi(action.Tags.Get("index")); err != nil {
		return
	} else {
		d.Content = markup.Code[sortedNames()[selectedIndex]]
		d.Update()
		ctx.Defer(func(context app.Context) {
			app.Window().Get("Prism").Call("highlightAll")
		})
	}
}
