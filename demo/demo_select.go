package demo

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/list"
	"github.com/mlctrez/goapp-mdc/pkg/selectm"
	"github.com/mlctrez/goapp-mdc/pkg/snackbar"
)

type SelectDemo struct {
	app.Compo
	base.JsUtil
	snack *snackbar.SnackBar
}

func buildList(hasEmpty bool, selected int) list.Items {
	items := list.Items{}
	if hasEmpty {
		items = append(items, &list.Item{Type: list.ItemTypeOption, Text: "", Value: ""})
	}

	for i := 0; i < 10; i++ {
		iStr := fmt.Sprintf("%d", i)
		items = append(items, &list.Item{Type: list.ItemTypeOption, Text: "Option " + iStr, Value: "value" + iStr})
	}
	items.Select(selected)
	return items
}

func (s *SelectDemo) Render() app.UI {

	if s.snack == nil {
		s.snack = &snackbar.SnackBar{ActionButtonText: "OK", LabelText: "messages will appear here"}
	}

	var rows []app.UI

	for _, outlined := range []bool{false, true} {
		for _, required := range []bool{false, true} {
			for _, hasEmpty := range []bool{true, false} {
				for _, selectedIndex := range []int{-1, 2} {
					label := fmt.Sprintf("Outlined:%t Required:%t hasEmpty=%t selected=%d", outlined, required, hasEmpty, selectedIndex)
					sel := &selectm.MDCSelect{Label: "Label", ItemsLabel: "Item Label", Id: s.UUID(),
						Outlined: outlined,
						Items:    buildList(hasEmpty, selectedIndex),
						Required: required,
					}
					rows = append(rows, GridRow(label, sel))
				}
			}
		}
	}

	return PageBody(s.snack, layout.Grid().Body(rows...))
}

func (s *SelectDemo) OnMount(ctx app.Context) {
	s.snack.TimeoutMs(-1)
	ctx.Handle(string(selectm.MDCSelectChange), func(context app.Context, action app.Action) {
		s.snack.LabelText = fmt.Sprintf("%v", action)
		s.snack.Update()
		ctx.NewActionWithValue(string(snackbar.Open), s.snack)
	})
}
