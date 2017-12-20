package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MastersAcademy/go-course-2017/homeworks/sergii.suprun_differz/homework_5/boxes"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: " + os.Args[0] + " <N> <X>")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("bad N")
		return
	}
	x, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("bad X")
		return
	}

	box, _ := boxes.NewMyBox(n, x)
	box.Generate()
	fmt.Println("New box generated: \n" + box.PrintWeight())

	box.Shake(1)
	fmt.Println("Shake in corner 1: \n" + box.PrintWeight())

	box.Shake(2)
	fmt.Println("Shake in corner 2: \n" + box.PrintWeight())

	box.Shake(3)
	fmt.Println("Shake in corner 3: \n" + box.PrintWeight())

	box.Shake(4)
	fmt.Println("Shake in corner 4: \n" + box.PrintWeight())
}
