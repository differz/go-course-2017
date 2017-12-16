package main

import (
	"fmt"
)

func main() {
	height := 4

	if height < 2 {
		fmt.Println("Height value is invalid. Height must be at least 2.")
		return
	}

	fmt.Println("Height: ", height)

	for i := 0; i < height; i++ {
		firstHalf := ""
		for j := 1; j <= height; j++ {
			if j >= height-i {
				firstHalf += "#"
			} else {
				firstHalf += " "
			}
		}

		secondHalf := []rune(firstHalf)
		for i, j := 0, len(secondHalf)-1; i < j; i, j = i+1, j-1 {
			secondHalf[i], secondHalf[j] = secondHalf[j], secondHalf[i]
		}
		fmt.Println(firstHalf + "  " + string(secondHalf))
	}

}
