package main

import "fmt"

func main() {
	const height = 4
	const X = "#"

	if height <= 2 {
		fmt.Println("The height must have a positive integer and >= 2")
		return
	}

	fmt.Println("Height:", height)

	for i, s := 0, X; i < height; i, s = i+1, s+X {
		fmt.Printf("%*[2]s  %[2]s\n", height, s)
	}
}