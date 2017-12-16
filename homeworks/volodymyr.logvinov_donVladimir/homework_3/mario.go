package main

import (
	"fmt"
	"os"
	"strconv"
)

const defaultHeight = 4
const innerGupSize = 2

func main() {
	h := defaultHeight

	args := os.Args[1:]

	if len(args) >= 1 {
		input, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Use \"mario number\" number should be a number")
			os.Exit(2)
		}

		h = input

		if h < 2 || h > 24 {
			fmt.Println("number should be in range from 2 to 24")
			os.Exit(3)
		}
	}

	fmt.Printf("Heigh: %d\n", h)

	for str := 1; str <= h; str++ {

		for col := 1; col <= 2*h+innerGupSize; col++ {
			switch {
			case col <= h && h-col < str:
				fmt.Print("#")
			case col > h+innerGupSize && col-h-innerGupSize <= str:
				fmt.Print("#")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")

	}
}
