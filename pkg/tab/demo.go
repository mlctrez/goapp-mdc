package tab

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Demo struct {
	app.Compo
}

func (d *Demo) Render() app.UI {
	id := uuid.New().String()

	bar := NewBar(id, []*Tab{
		NewTab("Tab One", 0).Active(),
		NewTab("Tab Two", 1).Icon("api"),
		NewTab("Tab Three", 2).Icon("favorite"),
	})
	bar.ActivateCallback(func(index int) {
		fmt.Println("you clicked on tab index", index)
	})

	return bar
}
