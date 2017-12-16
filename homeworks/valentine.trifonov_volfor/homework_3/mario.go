package main

import (
	"fmt"
	"strings"
)

var height = 4
var symbol = "#"

func main() {
	fmt.Println("Height:", height)

	for i := 1; i <= height; i++ {
		s := strings.Repeat(symbol, i)
		fmt.Printf("%s%s  %s\n", strings.Repeat(" ", height-i), s, s)
	}
}
