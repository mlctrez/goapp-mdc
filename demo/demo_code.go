package demo

import (
	"fmt"
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
	list    *list.List
}

func (d *CodeDemo) OnNav(ctx app.Context) {
	d.SetActive(urlFragmentToInt(ctx))
	d.list.Select(d.Active)
	d.Update()
	ctx.Defer(prismHiglightAll)
}

func (d *CodeDemo) SetActive(index int) {
	if index > 0 && index < len(markup.Code) {
		d.Active = index
	} else {
		d.Active = 0
	}
}

func urlFragmentToInt(ctx app.Context) (result int) {
	for i, code := range markup.Code {
		if code.Name == ctx.Page().URL().Fragment {
			return i
		}
	}
	return
}

func (d *CodeDemo) Render() app.UI {
	if d.list == nil {
		items := list.Items{}
		for _, c := range markup.Code {
			items = append(items, &list.Item{Text: c.Name,
				Type: list.ItemTypeAnchor,
				Href: fmt.Sprintf("/code#%s", c.Name),
			})
		}
		d.list = &list.List{Id: "codeNav", Type: list.Navigation, Items: items.UIList()}
	}

	body := &drawer.Drawer{Id: "codeNavigation", Type: drawer.Standard, List: d.list}
	d.Content = string(markup.Decode(markup.Code[d.Active].Code))
	return PageBody(body, app.Raw(d.Content))
}

func prismHiglightAll(_ app.Context) {
	prism := app.Window().Get("Prism")
	if prism.Truthy() {
		prism.Call("highlightAll")
	}
}

func (d *CodeDemo) OnMount(ctx app.Context) {
	ctx.Defer(prismHiglightAll)
}

func (d *CodeDemo) eventHandler(ctx app.Context, action app.Action) {
	if selectedIndex, err := strconv.Atoi(action.Tags.Get("index")); err != nil {
		return
	} else {
		if selectedIndex > -1 && selectedIndex < len(markup.Code) {
			ctx.Navigate(fmt.Sprintf("/code#%s", markup.Code[selectedIndex].Name))
		}
	}
}
