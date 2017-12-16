package main

import "fmt"

const N = 4

var n1, n2 int = N - 1, 0

func main() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j < n1 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("  ")
		for k := 0; k < N; k++ {
			if k > n2 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		n1--
		n2++
		fmt.Println()
	}
}
