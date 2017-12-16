package main

import (
	"fmt"
	"strings"
)

var numrows int

const minrows = 2
const defaultrows = 4

func main() {
	fmt.Print("Enter Mario number: ")
	fmt.Scanln(&numrows)

	if numrows < minrows {
		fmt.Println("Too small value, set to default: ", defaultrows)
		numrows = defaultrows
	}

	fmt.Println()
	fmt.Println("Height of the pyramide: ", numrows)
	fmt.Println()
	fmt.Println("Cheater code:")

	// Cheater code
	for i := 1; i <= numrows; i++ {
		d := strings.Repeat(" ", numrows-i)
		fmt.Print(d)
		d = strings.Repeat("#", i)
		fmt.Print(d + "  ")
		fmt.Println(d)
	}

	fmt.Println()
	fmt.Println("Not aesthetic but working code:")

	// Not aesthetic but working code
	for i := 1; i <= numrows; i++ {
		for j := numrows - i; j >= 1; j-- {
			fmt.Print(" ")
		}
		for n := 0; n < i; n++ {
			fmt.Print("#")
		}
		fmt.Print("  ")
		for n := 0; n < i; n++ {
			fmt.Print("#")
		}
		for j := numrows - i; j >= 1; j-- {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
