package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	place := "alexandr.kravchenko_drhihi"

	if len(os.Args) > 1 {
		place = strings.Join(os.Args[1:], " ")
	}

	fmt.Printf("Hello World from %s.\n", place)
}
