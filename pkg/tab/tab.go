package tab

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
)

type Tab struct {
	app.Compo
	Index  int
	Active bool
	Label  string
	Icon   icon.MaterialIcon
}

type Tabs []*Tab

func (ts Tabs) Select(index int) {
	for i, tab := range ts {
		tab.Index = i
		if index == i {
			tab.Active = true
		}
	}
}

func (ts Tabs) UIList() (body []app.UI) {
	for _, tab := range ts {
		body = append(body, tab)
	}
	return
}

func (t *Tab) iconUI() app.UI {
	var icon app.HTMLSpan
	if t.Icon != "" {
		icon = t.Icon.Span()
		icon.Class("mdc-tab__icon")
		icon.Aria("hidden", "true")
	}
	return icon
}

func (t *Tab) labelUI() app.UI {
	return app.Span().Class("mdc-tab__text-label").Text(t.Label)
}

func (t *Tab) indicatorUI() app.UI {
	root := app.Span().Class("mdc-tab-indicator")
	if t.Active {
		root.Class("mdc-tab-indicator--active")
	}
	indicator := app.Span()
	indicator.Class("mdc-tab-indicator__content mdc-tab-indicator__content--underline")
	return root.Body(indicator)
}

func (t *Tab) Render() app.UI {
	root := app.Button().Attr("role", "tab")
	root.Class("mdc-tab")

	if t.Active {
		root.Class("mdc-tab--active").Aria("selected", "true")
		root.Attr("tabindex", "0")
	} else {
		root.Attr("tabindex", "-1")
	}

	root.Body(
		app.Span().Class("mdc-tab__content").Body(t.iconUI(), t.labelUI()),
		t.indicatorUI(),
		app.Span().Class("mdc-tab__ripple"),
	)
	return root
}
