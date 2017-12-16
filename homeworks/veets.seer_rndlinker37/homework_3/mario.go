package main

import "fmt"

func main() {

	mario := 4 // default value
	X := "#"   // default char to print

	mariomin := 2
	mariomax := 24

	fmt.Print("Enter Mario number: ")
	fmt.Scanln(&mario) // input value

	if mario < mariomin {
		fmt.Println("Mario is too small, set to ", mariomin)
		mario = mariomin
	}
	if mario > mariomax {
		fmt.Println("Mario is too big, set to ", mariomax)
		mario = mariomax
	}

	fmt.Println("")

	for n := 1; n <= mario-3; n++ {
		fmt.Print(" ")
	}
	fmt.Println("Height: ", mario)
	for i := 0; i <= mario; i++ {
		fmt.Print(" ")
		for n := 1; n <= mario-i; n++ {
			fmt.Print(" ")
		}
		for n := 1; n <= i; n++ {
			fmt.Print(X)
		}
		fmt.Print("  ")
		for n := 1; n <= i; n++ {
			fmt.Print(X)
		}
		for n := 1; n <= mario-i; n++ {
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}

}
