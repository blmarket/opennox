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

func (l *ListElement[T, P]) Init_425770() *ListElement[T, P] {
	l.next = l
	l.prev = l
	l.head = nil
	return l
}

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
