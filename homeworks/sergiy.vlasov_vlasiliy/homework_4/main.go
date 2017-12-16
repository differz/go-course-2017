package main

import (
	"fmt"
	"errors"
)

type Item struct {
	value int
	pointer *Item
}

func main() {
	fmt.Println("-- Create list --")

	firstPointer := &Item{1, nil}

	currentPointer := firstPointer
	for i:=2; i <=10; i++ {
		//add to end
		currentPointer = addNext(currentPointer, i)
	}

	renderList(firstPointer)

	fmt.Println("----------------")

	fmt.Println("-- Make first --")
	//make first
	new := Item{0, firstPointer}
	renderList(&new)

	fmt.Println("----------------")

	fmt.Println("---- Insert ----")
	//insert after 4
	if _, err := insertElementAfter(4, &new); err != nil {
		fmt.Println(err.Error())
	} else {
		renderList(&new)
	}

	fmt.Println("----------------")

	fmt.Println("--- To count ---")
	fmt.Println("Count elemets:", countElements(&new))

	fmt.Println("----------------")

	fmt.Println("-- Delete element --")
	//delete 777
	if _, err := deleteElement(777, &new); err != nil {
		fmt.Println(err.Error())
	} else {
		renderList(&new)
	}

}

func addNext(item *Item, value int) *Item {
	next := &Item{value, nil}
	item.pointer = next

	return next
}

func renderList(pointer *Item) {
	currentPointer := pointer
	fmt.Println(*currentPointer)

	if (pointer.pointer != nil) {
		renderList(pointer.pointer)
	}
}

func insertElementAfter(value int, firstPointer *Item) (bool, error) {
	currentPointer := firstPointer

	for {
		if currentPointer.pointer == nil {
			return false, errors.New("Insert element: value is not found")
		}

		if currentPointer.value == value {
			pointer := currentPointer.pointer
			new := Item{777, pointer}
			currentPointer.pointer = &new

			return true, nil
		}

		currentPointer = currentPointer.pointer
	}
}

func deleteElement(value int, firstPointer *Item) (bool, error) {
	currentPointer := firstPointer

	for {
		if currentPointer.value != value && currentPointer.pointer == nil {
			return false, errors.New("Delete element: value is not found")
		}

		if currentPointer.pointer.value == value {
			if currentPointer.pointer.pointer == nil {
				currentPointer.pointer = nil
			} else {
				currentPointer.pointer = currentPointer.pointer.pointer
			}

			return true, nil

		}

		currentPointer = currentPointer.pointer
	}
}

func countElements(firstPointer *Item) int {
	var count int

	currentPointer := firstPointer

	for {
		count++
		if currentPointer.pointer == nil {
			break
		}
		currentPointer = currentPointer.pointer
	}

	return count
}