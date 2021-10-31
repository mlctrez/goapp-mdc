package list

import "testing"

func TestItems_Select(t *testing.T) {

	var items Items
	// should not blow up
	items.Select(0).Select(0)

	index := -1 // tab 0 on first item only
	items = append(items, &Item{}, &Item{}, &Item{})
	if items.Select(index)[0].state != ItemSelectStateTabZero {
		t.Error("items[0].state != ItemSelectStateTabZero")
	}
	if items.Select(index)[1].state != ItemSelectStateNone {
		t.Error("items[1].state != ItemSelectStateNone")
	}
	index = 0 // first selected, rest not selected
	if items.Select(index)[0].state != ItemSelectStateSelected {
		t.Error("items[0].state != ItemSelectStateTabZero")
	}
	if items.Select(index)[1].state != ItemSelectStateNotSelected {
		t.Error("items[1].state != ItemSelectStateNone")
	}

}
