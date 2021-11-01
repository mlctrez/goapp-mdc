package demo

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/tab"
)

type TabDemo struct {
	app.Compo
}

func (d *TabDemo) Render() app.UI {

	id := uuid.New().String()

	tab := tab.NewBar(id, []*tab.Tab{
		tab.NewTab("Tab One", 0).Active(),
		tab.NewTab("Tab Two", 1).Icon("api"),
		tab.NewTab("Tab Three", 2).Icon("favorite"),
	})
	tab.ActivateCallback(func(index int) {
		fmt.Println("you clicked on tab index", index)
	})

	return PageBody(tab)
}
