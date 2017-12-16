package LinkedList

import (
	"fmt"
	"testing"
	"log"
)

var l *LinkedList

func prinList(l *LinkedList) {
	fmt.Print("linked list: ")
	for k, v := range l.Value() {
		fmt.Printf("{%d: %v}", k, v)
	}
	fmt.Println()
}

func initLinkedList() {
	fmt.Println("--init--")
	l = nil
	l = Append(l,"one")
	l = Append(l,2)
	l = Append(l,"three")
	l = Append(l,4)
	fmt.Printf("len: %d\n", Len(l))
	prinList(l)
	fmt.Printf("----\n\n")
}

func TestAppend(t *testing.T) {
	defer fmt.Println()
	log.Println("======= TestAppend =======")
	initLinkedList()
}

func TestGet(t *testing.T) {
	defer fmt.Println()
	log.Println("======= TestGet =======")
	initLinkedList()
	f := func(i uint) {
		v, err := l.Get(i)
		fmt.Printf("get %d: %v { err: %v }\n", i, v, err)
	}
	f(1)
	f(5)
}

func TestSet(t *testing.T) {
	defer fmt.Println()
	log.Println("======= TestSet =======")
	initLinkedList()
	f := func(i uint, v interface{}) {
		err := l.Set(i, v)
		fmt.Printf("set %d: %v { err: %v }\n", i, v, err)
		prinList(l)
	}
	f(1, "two")
	f(5, "five")
}

func TestDelete(t *testing.T) {
	defer fmt.Println()
	log.Println("======= TestDelete =======")
	initLinkedList()
	f := func(i uint) {
		err := l.Delete(i)
		fmt.Printf("del %d { err: %v }\n", i, err)
		prinList(l)
	}
	f(5)
	f(1)
}
