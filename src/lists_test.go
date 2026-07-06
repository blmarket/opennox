package opennox

import (
	"testing"
	"unsafe"
)

func TestListHeadClear(t *testing.T) {
	var l listHead[listItem, *listItem]
	l.Clear()
	if l.next != &l.listItem || l.prev != &l.listItem || l.head != &l.listItem {
		t.Error("Clear did not initialize list properly")
	}
}

func TestListHeadFirst(t *testing.T) {
	var l listHead[listItem, *listItem]
	l.Clear()
	if l.First() != nil {
		t.Error("First on empty list should return nil")
	}
}

func TestListItemNext(t *testing.T) {
	var item listItem
	if item.Next() != nil {
		t.Error("Next on uninitialized item should return nil")
	}

	var l listHead[listItem, *listItem]
	l.Clear()
	item2 := &listItem{}
	l.Append(item2)

	// After appending, First should return the item
	first := l.First()
	if first == nil {
		t.Error("First should not be nil after Append")
	}
}

func TestListItemRemove(t *testing.T) {
	var l listHead[listItem, *listItem]
	l.Clear()
	item := &listItem{}
	l.Append(item)

	item.Remove()
	// After removal, list should be empty
	if l.First() != nil {
		t.Error("First should be nil after removing the only item")
	}
}

func TestNoxCommonListFunctions(t *testing.T) {
	var l listHead[listItem, *listItem]
	l.Clear()

	// Test getFirstSafe
	ptr := nox_common_list_getFirstSafe_425890(unsafe.Pointer(&l))
	if ptr != nil {
		t.Error("getFirstSafe on empty list should return nil")
	}

	// Test getNextSafe with nil
	ptr = nox_common_list_getNextSafe_4258A0(nil)
	if ptr != nil {
		t.Error("getNextSafe with nil should return nil")
	}

	// Test clear
	nox_common_list_clear_425760(unsafe.Pointer(&l))

	// Test append
	item := &listItem{}
	nox_common_list_append_4258E0(unsafe.Pointer(&l), unsafe.Pointer(item))

	// Test getNext
	ptr = nox_common_list_getNext_425940(unsafe.Pointer(item))
	// Should return nil since it's the only item and next points to head
}
