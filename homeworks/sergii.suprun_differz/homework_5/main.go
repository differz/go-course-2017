package main

import (
	"fmt"
	"log"
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

	box, err := boxes.NewMyBox(n, x)
	if err != nil {
		log.Panicf("Cant create box")
	}

	box.Generate()
	fmt.Println("New box generated: \n" + box.PrintWeight())

	box.Shake(boxes.ByCorner1)
	fmt.Println("Shake in corner 1: \n" + box.PrintWeight())

	box.Shake(boxes.ByCorner2)
	fmt.Println("Shake in corner 2: \n" + box.PrintWeight())

	box.Shake(boxes.ByCorner3)
	fmt.Println("Shake in corner 3: \n" + box.PrintWeight())

	box.Shake(boxes.ByCorner4)
	fmt.Println("Shake in corner 4: \n" + box.PrintWeight())
}
