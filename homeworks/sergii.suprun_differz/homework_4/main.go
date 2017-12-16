package main

import "fmt"

import "github.com/differz/go-course-2017/homeworks/sergii.suprun_differz/homework_4/linkedlist"
import "github.com/differz/go-course-2017/homeworks/sergii.suprun_differz/homework_4/list"

func main() {

	var list list.List
	list = linkedlist.New()

	list.Add("element 0")
	list.Add("element 1")
	list.Add("element 2")
	list.Add("element 3")

	for i := 0; i <= list.Size(); i++ {
		element, error := list.Get(i)
		if error != nil {
			fmt.Println(error)
			break
		}
		fmt.Println(element)
	}

	fmt.Println(list.String())

	list.Remove("element 1")
	fmt.Println(list.String())

	list.Remove(1)
	fmt.Println(list.String())

	list.Set(1, "last element")
	fmt.Println(list)

	list.Remove(0)
	fmt.Println(list.String())

	list.Insert(0, "ins 0")
	list.Insert(1, "ins 1")
	list.Insert(4, "ins 4")
	fmt.Println(list.String())

	list.Remove(nil)
	fmt.Println(list.String())

}
