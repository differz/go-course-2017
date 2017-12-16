package main

import "fmt"

func main() {

	for {
		var pyramid_height int
		fmt.Print("Enter height >= 2 : ")
		fmt.Scanf("%d", &pyramid_height)

		if pyramid_height >= 2 {
			fmt.Printf("Heigh: %d\n", pyramid_height)

			for i := 0; i < pyramid_height; i++ {
				for j := 0; j < pyramid_height-i; j++ {
					fmt.Printf("%s", " ")
				}
				for k := 0; k < i+1; k++ {
					fmt.Printf("#")
				}
				fmt.Printf(" ")
				for k := 0; k < i+1; k++ {
					fmt.Printf("#")
				}
				fmt.Printf("\n")
			}
		}
	}
}
