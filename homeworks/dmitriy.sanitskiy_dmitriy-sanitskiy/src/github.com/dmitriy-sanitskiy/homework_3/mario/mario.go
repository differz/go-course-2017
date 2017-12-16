package main

import "fmt"

func main() {
	var h int
	fmt.Println("Enter pyramid's height")
	fmt.Scan(&h)
	fmt.Println("Height: ", h)

	for i := 0; i < h; i++ {
		for s := 0; s < h-i; s++ {
			fmt.Print(" ")
		}

		for x := h - i; x <= h; x++ {
			fmt.Print("#")
		}
		fmt.Print("  ")

		for x := h - i; x <= h; x++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
