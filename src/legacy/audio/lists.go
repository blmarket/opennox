package audio

import "unsafe"

type ListItem struct {
	next *ListItem
	prev *ListItem
	head *ListItem
}

type ListElement[T any] struct {
	next *ListElement[T]
	prev *ListElement[T]
	head *ListElement[T]
}

func (l *ListElement[T]) PromoteUnsafe() *T {
	return (*T)(unsafe.Pointer(l))
}

func (l *ListElement[T]) IsHead() bool {
	return l.head == l
}

func (l *ListElement[T]) Next() *ListElement[T] {
	return l.next
}

// It's init for non-head element
func (l *ListElement[T]) Init_425770() *ListElement[T] {
	l.next = l
	l.prev = l
	l.head = nil
	return l
}

// It's init for head element
func (l *ListElement[T]) Clear_425760() {
	l.next = l
	l.prev = l
	l.head = l
}

func (l *ListElement[T]) NextSafe_425940() *T {
	it := l.next
	if it != nil && it == it.head {
		return nil
	}
	return it.PromoteUnsafe()
}

func (l *ListElement[T]) NextSafe_4258A0() *T {
	if l == nil {
		return nil
	}
	return l.NextSafe_425940()
}

func (l *ListElement[T]) FirstSafe_425890() *T {
	return l.FirstSafe_4258A0()
}

func (l *ListElement[T]) FirstSafe_4258A0() *T {
	if l == nil {
		return nil
	}
	return l.NextSafe_425940()
}

func (a1p *ListElement[T]) PrevSafe_425960() *T {
	if a1p.prev.head != a1p.prev {
		return a1p.prev.PromoteUnsafe()
	}
	return nil
}

func (list *ListElement[T]) Append_4258E0(cur *ListElement[T]) {
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

func (e *ListElement[T]) Remove_425920() {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = e
	e.prev = e
}
