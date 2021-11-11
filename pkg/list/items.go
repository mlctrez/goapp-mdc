package list

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Items allows setting up the state of a list of items before the first render. Setting the tab index
// or pre-selected item and associated attributes of the items in the list.
type Items []*Item

// SelectedItemText returns the text of the currently selected item or empty if no selection made.
//
// This should be called after Select()
func (items Items) SelectedItemText() (result string) {
	for _, item := range items {
		if item.state == ItemSelectStateSelected {
			result = item.Text
		}
	}
	return
}

func (items Items) UIList() (uis []app.UI) {
	for _, item := range items {
		uis = append(uis, item)
	}
	return
}

func (items Items) Select(index int) Items {
	if items == nil {
		return items
	}
	for i, item := range items {
		item.state = determineSelectState(i, index)
	}
	return items
}

func (items Items) SelectHref(href string) Items {
	if items == nil {
		return items
	}
	index := -1
	for i, item := range items {
		if item.Href == href {
			index = i
		}
	}
	return items.Select(index)
}

func determineSelectState(listIndex, selectedIndex int) ItemSelectState {
	if selectedIndex < 0 {
		if listIndex == 0 {
			return ItemSelectStateTabZero
		}
		return ItemSelectStateNone
	} else {
		if listIndex == selectedIndex {
			return ItemSelectStateSelected
		}
		return ItemSelectStateNotSelected
	}
}

type ItemSelectState int

const (
	ItemSelectStateNone ItemSelectState = iota
	ItemSelectStateTabZero
	ItemSelectStateSelected
	ItemSelectStateNotSelected
)
