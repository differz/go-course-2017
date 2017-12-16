package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 4 {
		panic("fake parameters")
	}

	box := initBox(os.Args[1:]...)
	fillBox(box)

	fmt.Println("")
	box.PrintlnDetails()
	fmt.Printf("\noriginal box\n")
	box.Println()

	box.sortCorner(1)
	fmt.Printf("\ncorner #1\n")
	box.Println()

	box.sortCorner(2)
	fmt.Printf("\ncorner #2\n")
	box.Println()

	box.sortCorner(3)
	fmt.Printf("\ncorner #3\n")
	box.Println()

	box.sortCorner(4)
	fmt.Printf("\ncorner #4\n")
	box.Println()

}
