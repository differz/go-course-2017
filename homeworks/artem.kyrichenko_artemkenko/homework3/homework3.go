package main

import "fmt"

func main() {
	Height :=10
	fmt.Print("Hight is ", Height, "\n")
	if Height >= 2 {
		for i := 1; i != Height; i++  {
			for j := Height; j != i - 1; j--  {
				fmt.Print(" ")
			}
			for k := 1; k != i; k++ {
				fmt.Print("#")
			}
			fmt.Print("  ")
			for k := 1; k != i ; k++  {
				fmt.Print("#")
			}
			fmt.Print("\n")
		}
	}

	return
}