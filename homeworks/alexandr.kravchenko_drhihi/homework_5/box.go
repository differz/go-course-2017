package main

import (
	"fmt"
	"github.com/drhihi/go-course-2017/homeworks/alexandr.kravchenko_drhihi/homework_5/models"
	"math/rand"
	"sort"
	"strconv"
)

type Box struct {
	a, b, height int
	sell         []models.Shape
}

const maxSizeShape = 10

/*
param:
	1 2
	3 4
*/
func (b *Box) sortCorner(corner int) {

	slice := b.getSliceSell()
	sort.Sort(slice)

	switch corner {
	default:
		panic("the corner is not right")

	case 1:
		for k, y, x, fy, fx := 0, 0, 0, 0, 0; k < len(slice); k++ {
			b.sell[y][x] = slice[k]
			y--
			x++
			if y < 0 || x > b.height-1 {
				fy++
				if fy > b.height-1 {
					fy--
					fx++
				}
				y = fy
				x = fx
			}
		}

	case 2:
		for k, y, x, fy, fx := 0, 0, b.height-1, 0, b.height-1; k < len(slice); k++ {
			b.sell[y][x] = slice[k]
			y++
			x++
			if y > b.height-1 || x > b.height-1 {
				fx--
				if fx < 0 {
					fx++
					fy++
				}
				y = fy
				x = fx
			}
		}

	case 3:
		for k, y, x, fy, fx := 0, b.height-1, 0, b.height-1, 0; k < len(slice); k++ {
			b.sell[y][x] = slice[k]
			y--
			x--
			if y < 0 || x < 0 {
				fx++
				if fx > b.height-1 {
					fx--
					fy--
				}
				y = fy
				x = fx
			}
		}

	case 4:
		for k, y, x, fy, fx := 0, b.height-1, b.height-1, b.height-1, b.height-1; k < len(slice); k++ {
			b.sell[y][x] = slice[k]
			y++
			x--
			if y > b.height-1 || x < 0 {
				fy--
				if fy < 0 {
					fy++
					fx--
				}
				y = fy
				x = fx
			}
		}

	}

}

func (b *Box) Println() {
	for _, vl := range b.sell {
		fmt.Println(vl)
	}
}

func (b *Box) PrintlnDetails() {
	for _, vl := range b.sell {
		fmt.Println("{", vl)
		for _, vr := range vl {
			fmt.Printf("%T %v\n", vr, vr)
		}
		fmt.Println("}")
	}
}

func initBox(param ...string) *Box {

	box := new(Box)

	for k, v := range param {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		switch k {
		case 0:
			box.a = i
		case 1:
			box.b = i
		case 2:
			box.height = i
		}
	}

	box.sell = make([]models.Shape, box.height)

	for i := 0; i < box.height; i++ {
		box.sell[i] = make(models.Shape, box.height)
	}

	return box
}

func fillBox(box *Box) {
	types := [...]string{
		"cube",
		"cone",
		"piramid",
	}

	for _, vl := range box.sell {
		for k, _ := range vl {
			switch types[rand.Intn(len(types))] {
			default:
				panic("unknown type")
			case "cube":
				vl[k] = models.NewCube(
					rand.Intn(maxSizeShape), // a
				)
			case "cone":
				vl[k] = models.NewCone(
					rand.Intn(maxSizeShape), // r
					rand.Intn(maxSizeShape), // h
				)
			case "piramid":
				vl[k] = models.NewPiramid(
					rand.Intn(maxSizeShape), // a
					rand.Intn(maxSizeShape), // b
					rand.Intn(maxSizeShape), // h
				)
			}
		}
	}

}

func (b *Box) getSliceSell() models.Shape {
	result := make(models.Shape, b.height*b.height)
	var i int
	for _, vl := range b.sell {
		for _, vr := range vl {
			result[i] = vr
			i++
		}
	}
	return result
}
