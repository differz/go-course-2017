package main

import "fmt"

var (
	height = 8
)

func printSpaces(n int) {
	for j := (height - n); j > 0; j-- {
		fmt.Printf(" ")
	}
}

func printBlocks(n int) {
	for j := 1; j <= n; j++ {
		fmt.Printf("#")
	}
}

func main() {
	fmt.Printf("Height: %v\n", height)
	for n := 1; n <= height; n++ {
		fmt.Printf("|")
		printSpaces(n)
		printBlocks(n)
		fmt.Printf("  ")
		printBlocks(n)
		printSpaces(n)
		fmt.Printf("|")
		fmt.Printf("\n")
	}
}
