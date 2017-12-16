package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var height = 4
	if len(os.Args) > 1 {
		parsedHeight, err := strconv.ParseInt(os.Args[1], 10, 0)
		if err == nil {
			height = int(parsedHeight)
		}
	}

	symbol := "#"

	if height%2 != 0 {
		symbol = "💩"
	}

	for i := 0; i < height; i++ {
		for j := 0; j < height-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(symbol)
		}
		fmt.Print("  ")
		for j := 0; j <= i; j++ {
			fmt.Print(symbol)
		}
		fmt.Print("\n")
	}
}
