package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	var simbol = "#"
	var cnt = 4

	//get command line arguments
	args := os.Args[1:]

	//check argument 1
	if len(args) >= 1 && len(args[0]) == 1 {
		simbol = args[0]
	}

	//check argument 2
	if len(args) == 2 {
		count, err := strconv.Atoi(args[1])
		if err != nil || count < 2{
			fmt.Println(err)
			return
		} else {
			cnt = count
		}
	}

	//build piramides
	for i := 1; i <= cnt; i++ {
		fmt.Println(
			concatIdenticalChars(" ", cnt - i) +
			concatIdenticalChars(simbol, i) +
			"  "+
			concatIdenticalChars(simbol, i)	+
			concatIdenticalChars(" ", cnt - i))
	}
}

func concatIdenticalChars(simbol string, count int) string {
	var result = ""
	for i := 1; i <= count; i++ {
		result += simbol
	}

	return result
}




