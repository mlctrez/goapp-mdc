package older

import (
	"sort"

	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/mlctrez/goapp-mdc/pkg/tab"
)

func buildDemoMap() map[string]app.UI {
	return map[string]app.UI{
		//"banner":   &demo.BannerDemo{},
		//"button":   &button.Demo{},
		//"card":     &card.Demo{},
		//"icon":     &icon.Demo{},
		//"drawer":   &drawer.Demo{},
		//"list":     &list.Demo{},
		//"checkbox": &checkbox.Demo{},
		//"dialog":   &dialog.Demo{},
		//"form":     &example.Example{},
		//"fab":      &fab.Demo{},
	}
}

type Demo struct {
	app.Compo
	bar         *tab.Bar
	ActiveIndex int
	DemoMap     map[string]app.UI
}

func (d *Demo) sortedDemoNames() []string {
	var sortedNames []string
	for k := range d.DemoMap {
		sortedNames = append(sortedNames, k)
	}
	sort.Strings(sortedNames)
	return sortedNames
}

func (d *Demo) Render() app.UI {

	if d.DemoMap == nil {
		d.DemoMap = buildDemoMap()
		var tabs []*tab.Tab
		var index int
		for _, groupName := range d.sortedDemoNames() {
			newTab := tab.NewTab(groupName, index)
			if d.ActiveIndex == index {
				newTab.Active()
			}
			tabs = append(tabs, newTab)
			index++
		}
		d.bar = tab.NewBar("demoPageBar", tabs)
		d.bar.ActivateCallback(func(index int) {
			d.ActiveIndex = index
			d.Update()
		})
	}

	return app.Div().Body(
		//&demo.UpdateBanner{},
		d.bar,
		d.DemoMap[d.sortedDemoNames()[d.ActiveIndex]],
	)
}
