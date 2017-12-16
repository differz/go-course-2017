package main

import "fmt"

func main() {
	var height int = 5

	for i := 1; i <= height; i++  {
		for index := height; index >= 0; index--  {
			if (index < i) {
				fmt.Printf(string('#'));
				continue;
			} else {
				fmt.Printf(string(' '));
			}
		}

		fmt.Printf(string(' '));

		for index := height; index >= 0; index--  {
			if (index < i) {
				fmt.Printf(string('#'));
				continue;
			}
		}

	fmt.Println(string(' '));
	}
}
