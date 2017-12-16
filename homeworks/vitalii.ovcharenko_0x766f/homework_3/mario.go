package main

import (
	"fmt"
)

func main() {
	const height = 4
	const block = "#"

	if height < 2 {
		fmt.Println("The height must have a positive integer and >= 2")
		return
	}

	fmt.Println("Height:", height)

	for i, s := 0, block; i < height; i, s = i+1, s+block {
		fmt.Printf("%*[2]s  %[2]s\n", height, s)
	}
}
