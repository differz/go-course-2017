package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MinH = 2
	MaxH = 23
	Sym  = "#"
	DefH = 4
)

func main() {

	h := DefH

	if len(os.Args) > 1 {
		var err error
		s := strings.Join(os.Args[1:], " ")
		h, err = strconv.Atoi(s)
		switch {
		case nil != err:
			fmt.Println("ERROR: argument isn't an integer.")
			os.Exit(1)
		case MinH > h || h > MaxH:
			fmt.Printf("ERROR: argument is invalid. Allowed range: %d...%d.\n", MinH, MaxH)
			os.Exit(2)
		}
	}

	fmt.Printf("Height: %d\n", h)

	if true {
		sf := "%" + strconv.Itoa(h) + "[1]s  %[1]s\n"
		for i := 1; i <= h; i++ {
			s := strings.Repeat(Sym, i)
			fmt.Printf(sf, s)
		}
	}

}
