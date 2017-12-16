package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 && (args[0] == "--help" || args[0] == "-h") {
		text :=
			"Just put this binary to your homework folder ant that's it.\n\n" +
			"Exit codes:\n" +
			"\t0: Ok\n" +
			"\t1: WTF?\n" +
			"\t2: Can't find folder with proper name: firstName.lastName_githubName in current path"

		fmt.Println(text)
		os.Exit(0)
	}

	path, err := os.Getwd()

	if err != nil {
		fmt.Println("Congrats!!! You got the error: ", err)
		os.Exit(1)
	}

	re := regexp.MustCompile("\\w+\\.\\w+_[\\w-]+")
	sub := re.FindString(path)

	if len(sub) == 0 {
		fmt.Println("Hey!!! Put me back")
		os.Exit(2)
	}

	fmt.Println(sub)
}