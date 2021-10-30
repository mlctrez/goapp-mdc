package fab

import (
	"sort"

	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/tab"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Demo struct {
	app.Compo
	activeIndex int
}

func (d *Demo) Render() app.UI {
	var tabs []*tab.Tab

	sortedGroupNames := sortedGroupNames()
	for index, groupName := range sortedGroupNames {
		newTab := tab.NewTab(groupName, index)
		if d.activeIndex == index {
			newTab.Active()
		}
		tabs = append(tabs, newTab)
	}

	bar := tab.NewBar("fabDemoTabBar", tabs)
	bar.ActivateCallback(func(index int) {
		d.activeIndex = index
		d.Update()
	})

	var content []app.UI
	content = append(content, bar)

	groupName := sortedGroupNames[d.activeIndex]
	iconNames := icon.AllGroupFunctions()[groupName]()
	for _, n := range iconNames {
		content = append(content, &Fab{Id: "fab_" + string(n), Icon: n})
	}

	return app.Div().Body(content...)
}

func sortedGroupNames() []string {
	var groupNamesSorted []string
	for groupName := range icon.AllGroupFunctions() {
		groupNamesSorted = append(groupNamesSorted, groupName)
	}
	sort.Strings(groupNamesSorted)
	return groupNamesSorted
}
