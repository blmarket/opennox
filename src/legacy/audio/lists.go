package audio

import "unsafe"

type ListItem struct {
	next *ListItem
	prev *ListItem
	head *ListItem
}

type ListElement[T any, P interface {
	*T
}] struct {
	next *ListElement[T, P]
	prev *ListElement[T, P]
	head *ListElement[T, P]
}

func (l *ListElement[T, P]) PromoteUnsafe() *T {
	return (*T)(unsafe.Pointer(l))
}

func (l *ListElement[T, P]) IsHead() bool {
	return l.head == l
}

func (l *ListElement[T, P]) Next() *ListElement[T, P] {
	return l.next
}

// It's init for non-head element
func (l *ListElement[T, P]) Init_425770() *ListElement[T, P] {
	l.next = l
	l.prev = l
	l.head = nil
	return l
}

// It's init for head element
func (l *ListElement[T, P]) Clear_425760() {
	l.next = l
	l.prev = l
	l.head = l
}

func (l *ListElement[T, P]) NextSafe_425940() *ListElement[T, P] {
	if l.next != nil && l.next == l.head {
		return nil
	}
	return l.next
}

func (l *ListElement[T, P]) FirstSafe_4258A0() *ListElement[T, P] {
	if l == nil {
		return nil
	}
	return l.NextSafe_425940()
}

func (list *ListElement[T, P]) Append_4258E0(cur *ListElement[T, P]) {
	if list == nil || cur == nil {
		panic("Append_4258E0 called will nil argument")
	}

	lastOrHead := list.prev

	if lastOrHead == nil && list.next == nil && list.head == nil {
		// List not initialized?!!
		list.Clear_425760()
		lastOrHead = list.prev
	}

	cur.next = list
	cur.prev = lastOrHead

	list.prev = cur
	if lastOrHead != nil {
		lastOrHead.next = cur
	}
}
