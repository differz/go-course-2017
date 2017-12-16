package models

import (
	"strconv"
)

type Cube struct {
	a      int
	weight int
}

/*
Params:
	side int
*/
func NewCube(param ...int) *Cube {
	c := new(Cube)
	for k, v := range param {
		switch {
		case 0 < k:
			panic("Error paraneters: " + strconv.Itoa(v))
		case 0 == k:
			c.SetA(v)
		}
	}
	return c
}

func (c *Cube) String() string { return strconv.Itoa(c.weight) }

func (c *Cube) getWeight() int { return c.weight }
func (c *Cube) setWeight()     { c.weight = c.a }

/*
side int
*/
func (c *Cube) GetA() int { return c.a }

/*
side int
*/
func (c *Cube) SetA(value int) {
	c.a = value
	c.setWeight()
}
