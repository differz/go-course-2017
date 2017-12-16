// https://golang.org/pkg/container/list/
// Package list implements a doubly linked list.
// To iterate over a list (where l is a *List):
// - подсчет размера списка;
// - добавление елемента в начало, конец и в середину списка;
// - удаление елемента из списка.
// rev 1.1 b
// todo - add func displayState(); seeker();

package main

import (
	"container/list"
	"fmt"
)

type element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next *element
	prev *element
	list *element // The list to which this element belongs.
	//	root  element     // sentinel list element, only &root, root.prev, and root.next are used
	len   int         // current list length excluding (this) sentinel element
	Value interface{} // The value stored with this element.
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.

func main() {

	L := list.New() //	var L list.List

	L.PushFront("1fins")
	L.PushFront("2fins")
	L.PushFront("3fins")
	L.PushFront("4fins")
	L.PushFront("5fins")

	L.PushBack("1bins")
	L.PushBack("2bins")
	L.PushBack("3bins")
	L.PushBack("4bins")
	// =====================================================================
	for e := L.Front(); e != nil; e = e.Next() { // Display state
		fmt.Print("->", e.Value, " ")
	}
	fmt.Println()

	fmt.Println("Len ", L.Len(), "First state") // Len() int - Len returns the number of elements of list l. The complexity is O(1).
	fmt.Println()
	// =====================================================================
	L.PushFront("frontinsert")
	L.PushBack("backinsert")
	// =====================================================================
	for e := L.Front(); e != nil; e = e.Next() { // Display state
		fmt.Print("->", e.Value, " ")
	}
	fmt.Println()

	fmt.Println("Len ", L.Len(), "First - last insert")
	fmt.Println()
	// =====================================================================
	middle := L.Len() / 2 // и в середину списка;

	e := L.Front()
	for ins := 0; ins < middle; ins++ { // seeker
		e = e.Next()
	}
	L.InsertBefore("middle", e) // inserter
	// =====================================================================
	for e := L.Front(); e != nil; e = e.Next() { // Display state
		fmt.Print("->", e.Value, " ")
	}
	fmt.Println()

	fmt.Println("Len ", L.Len(), "Middle insert to ", middle)
	fmt.Println()
	// =====================================================================
	any := 3 // remove # from 0 to last

	if any >= L.Len() {
		any = -1
		fmt.Println("# Error")
	} else {
		r := L.Front()
		for rem := 0; rem < any; rem++ { // search # to remove
			r = r.Next()
		}
		L.Remove(r) // # remover
	}
	// =====================================================================
	for e := L.Front(); e != nil; e = e.Next() { // Display state
		fmt.Print("->", e.Value, " ")
	}
	fmt.Println()

	fmt.Println("Len ", L.Len(), "Remove any #", any)
	fmt.Println()
}
