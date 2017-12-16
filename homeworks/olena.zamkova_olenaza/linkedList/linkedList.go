package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
)

type node struct {
	value int
	next *node
}

func getList() *node {
	var head *node

	fmt.Print("Enter list item or 'q' to exit: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		if scanner.Text() == "q" {break}
		data, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Print("Enter valid list item or 'q' to exit: ")
			continue
		}

		head = appendItem(head, data)

		fmt.Print("Enter list item or 'q' to exit: ")
	}

	return head
}

func checkInput (input string) (int, error) {
	data, err := strconv.Atoi(input)

	if err != nil {
		return 0, errors.New("input is invalid")
	}

	return data,nil
}

func printList(head *node) {
	fmt.Println("Your list is:")

	if head == nil {
		fmt.Println("empty")
		return
	}

	fmt.Printf("%d", head.value)
	curr := head

	for curr.next != nil {
		curr = curr.next
		fmt.Printf("  %d", curr.value)
	}

	fmt.Printf("\n")
	return
}

func count(head *node) int {
	if head == nil {
		return 0
	}

	curr := head
	n := 1

	for curr.next != nil {
		curr = curr.next
		n++
	}

	return n
}

func prependItem(head *node, data int) *node{
	var newNode node

	newNode.next = head
	newNode.value = data

	return &newNode
}

func appendItem(head *node, data int) *node {
	if head == nil {
		return prependItem(head, data)
	}

	curr := head

	for curr.next != nil {
		curr = curr.next
	}

	var appNode node
	appNode.value = data
	appNode.next = nil
	curr.next = &appNode

	return head
}

func insertItem(head *node, data int, pos int) (*node, error) {
	if pos == 1 {
		return prependItem(head, data), nil
	}

	if head == nil {
		return head, errors.New(fmt.Sprintf("Your list is empty. There is no (%pos)th position", pos))
	}

	curr := head

	for i := 1; i < pos - 1; i++ {
		curr = curr.next
		if curr.next == nil {
			return head, errors.New("your list is not long enough")
		}
	}

	var insNode node
	insNode.value = data
	insNode.next = curr.next
	curr.next = &insNode
	return head, nil
}

func deleteItem(head *node, data int) (*node, error) {
	err := errors.New("there is no data to delete")

	if head == nil {
		return head, err
	}

	if head.value == data {
		return head.next, nil
	}

	curr := head
	prev := head

	for curr.next != nil {
		curr = curr.next
		if curr.value == data {
			prev.next = curr.next
			return head, nil
		}
		prev = curr
	}

	return head, err
}

func main() {
	var item, pos int
	var input string
	var err error

	head := getList()

	printList(head)

	fmt.Printf("Your list contains %d items.\n", count(head))

	fmt.Println("Enter item to prepend: ")
	fmt.Scanln(&input)
	item, err = checkInput(input)
	if err != nil {
		fmt.Println(err)
	} else {
		head = prependItem(head, item)
	}

	printList(head)

	fmt.Println("Enter item to append: ")
	fmt.Scanln(&input)
	item, err = checkInput(input)
	if err != nil {
		fmt.Println(err)
	} else {
		head = appendItem(head, item)
	}

	printList(head)

	fmt.Println("Enter item to insert: ")
	fmt.Scanln(&input)
	item, err = checkInput(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Enter position to insert into: ")
		fmt.Scanln(&input)
		item, err = checkInput(input)
		if err != nil {
			fmt.Println(err)
		} else {
			head, err = insertItem(head, item, pos)

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	printList(head)

	fmt.Println("Enter item to delete: ")
	fmt.Scanln(&input)
	item, err = checkInput(input)
	if err != nil {
		fmt.Println(err)
	} else {
		head, err = deleteItem(head, item)

		if err != nil {
			fmt.Println(err)
		}
	}

	printList(head)
}
