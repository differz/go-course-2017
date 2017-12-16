package main

import (
	"fmt"
)

// List structure
type List struct {
	prev  *List
	value interface{}
	next  *List
}

// GetHead switches to the first element of the list
func GetHead(l *List) *List {
	for l.prev != nil {
		l = l.prev
	}
	return l
}

// PrintList Prints the list from first to last element
func (l *List) PrintList() {
	l = GetHead(l)
	for l != nil {
		fmt.Println(l.value)
		l = l.next
	}
}

// ListSize returns number of elements in the list
func (l *List) ListSize() int {
	l = GetHead(l)
	size := 0

	for l != nil {
		size++
		l = l.next
	}

	return size
}

// InsertInMiddle add an element in the middle of the list
func (l *List) InsertInMiddle(value interface{}) {
	l = GetHead(l)
	listSize := l.ListSize()
	middleIndex := listSize / 2
	for i := 0; i <= middleIndex; i++ {
		l = l.next
	}
	newElement := List{value: value, next: l, prev: l.prev}
	l.prev.next = &newElement
	l.prev = &newElement
}

// Append adds the element to an end of list
func (l *List) Append(i interface{}) {
	for l.next != nil {
		l = l.next
	}

	l.next = &List{prev: l, value: i}
}

// Prepend add an element to the beginning of a list
func (l *List) Prepend(i interface{}) {
	l = GetHead(l)
	l.prev = &List{next: l, value: i}
}

// Delete removes element from list
func (l *List) Delete(i interface{}) {
	l = GetHead(l)
	for l != nil {
		if l.value == i {
			if l.prev != nil {
				l.prev.next = l.next
			}
			if l.next != nil {
				l.next.prev = l.prev
			}
			l = nil
			break
		}
		l = l.next
	}
}

func main() {

	list := List{value: 0}

	list.Prepend(2)
	list.Prepend(1)
	list.Append(5)
	list.Append(6)
	list.InsertInMiddle(4)

	list.Delete(1)
	list.PrintList()
	fmt.Println("LIST SIZE: ", list.ListSize())
}
