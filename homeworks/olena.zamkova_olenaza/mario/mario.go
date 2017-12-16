package main

import "fmt"

var height int = 4

func main() {
	if height >= 2 {
		fmt.Printf("Height: %d\n", height)

		for i := 1; i <= height; i++ {
			for j := 1; j <= height-i; j++ {
				fmt.Printf(" ")
			}
			for j := 1; j <= 2*i-1; j++ {
				fmt.Printf("#")
				if j == i {
					fmt.Printf("  ")
				}
			}
			fmt.Printf("#\n")
		}
	} else {
		fmt.Println("The height is invalid")
	}
}