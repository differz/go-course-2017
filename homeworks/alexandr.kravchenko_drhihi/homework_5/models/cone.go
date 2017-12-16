package models

import (
	"strconv"
)

type Cone struct {
	r      int
	h      int
	weight int
}

/*
Params:
	radius int
	height int
*/
func NewCone(params ...int) *Cone {
	c := new(Cone)
	for k, v := range params {
		switch {
		case 1 < k:
			panic("Error paraneters: " + strconv.Itoa(v))
		case 0 == k:
			c.SetR(v)
		case 1 == k:
			c.SetH(v)
		}
	}
	return c
}

func (c *Cone) String() string { return strconv.Itoa(c.weight) }

func (c *Cone) getWeight() int { return c.weight }
func (c *Cone) setWeight()     { c.weight = c.r }

/*
radius int
*/
func (c *Cone) GetR() int { return c.r }

/*
radius int
*/
func (c *Cone) SetR(value int) {
	c.r = value
	c.setWeight()
}

/*
height int
*/
func (c *Cone) GetH() int { return c.h }

/*
 height int
*/
func (c *Cone) SetH(value int) { c.h = value }
