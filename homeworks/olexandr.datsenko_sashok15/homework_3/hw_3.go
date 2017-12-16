package main

import "fmt"

var a int

func CreateStr(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "#"
	}
	return str
}

func main() {
	fmt.Print("Please, input a number ")
	fmt.Scan(&a)
	for true {
		if a >= 2 {
			break
		} else {
			fmt.Println("Fiasko. Try to input to larger number ")
			fmt.Scan(&a)
		}
	}

	fmt.Println("Height:", a)
	str := CreateStr(a)

	for i := 1; i <= a; i++ {
		for j := 1; j <= a-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println(str[:i], "", str[:i])
	}
}
