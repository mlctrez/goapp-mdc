package tab

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Tab struct {
	app.Compo
	index                int
	active               bool
	label                string
	icon                 string
}

func NewTab(label string, index int) *Tab {
	return &Tab{label: label, index: index}
}

func (t *Tab) Active() *Tab {
	t.active = true
	return t
}
func (t *Tab) Icon(text string) *Tab {
	t.icon = text
	return t
}

func (t *Tab) iconUI() app.UI {
	return app.If(t.icon != "",
		app.Span().Class("mdc-tab__icon material-icons").Aria(
			"hidden", "true").Text(t.icon),
	)
}

func (t *Tab) labelUI() app.UI {
	return app.Span().Class("mdc-tab__text-label").Text(t.label)
}

func (t *Tab) indicatorUI() app.UI {
	root := app.Span()
	if t.active {
		root.Class("mdc-tab-indicator mdc-tab-indicator--active")
	} else {
		root.Class("mdc-tab-indicator")
	}
	indicator := app.Span()
	indicator.Class("mdc-tab-indicator__content mdc-tab-indicator__content--underline")
	return root.Body(indicator)
}

func (t *Tab) Render() app.UI {

	root := app.Button().Attr("role", "tab")
	if t.active {
		root.Class("mdc-tab mdc-tab--active").Aria("selected", "true")
	} else {
		root.Class("mdc-tab")
	}
	root.Attr("tabindex", fmt.Sprintf("%d", t.index))

	root.Body(
		app.Span().Class("mdc-tab__content").Body(t.iconUI(), t.labelUI()),
		t.indicatorUI(),
		app.Span().Class("mdc-tab__ripple"),
	)
	return root
}

/*

   <button class="mdc-tab mdc-tab--active" role="tab" aria-selected="true" tabindex="0" id="mdc-tab-0">
       <span class="mdc-tab__content">
           <span class="mdc-tab__icon material-icons" aria-hidden="true">favorite</span>
           <span class="mdc-tab__text-label">Favorites</span>
       </span>
       <span class="mdc-tab-indicator mdc-tab-indicator--active">
           <span class="mdc-tab-indicator__content mdc-tab-indicator__content--underline"></span>
       </span>
       <span class="mdc-tab__ripple"></span>
   </button>
   <button class="mdc-tab" role="tab" tabindex="1" id="mdc-tab-1">
       <span class="mdc-tab__content">
           <span class="mdc-tab__icon material-icons" aria-hidden="true">bookmark</span>
           <span class="mdc-tab__text-label">Bookmark</span>
       </span>
       <span class="mdc-tab-indicator">
           <span class="mdc-tab-indicator__content mdc-tab-indicator__content--underline"></span>
       </span>
       <span class="mdc-tab__ripple"></span>
   </button>


*/
