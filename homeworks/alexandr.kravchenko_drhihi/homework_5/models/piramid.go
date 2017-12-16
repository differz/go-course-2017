package models

import "strconv"

type Piramid struct {
	a      int
	b      int
	h      int
	weight int
}

/*
Params:
	baseA int
	sideB int
	height int
*/
func NewPiramid(params ...int) *Piramid {
	p := new(Piramid)
	for k, v := range params {
		switch {
		case 2 < k:
			panic("Error paraneters: " + strconv.Itoa(v))
		case 0 == k:
			p.SetA(v)
		case 1 == k:
			p.SetB(v)
		case 2 == k:
			p.SetH(v)
		}
	}
	return p
}

func (p *Piramid) String() string { return strconv.Itoa(p.weight) }

func (p *Piramid) getWeight() int { return p.weight }
func (p *Piramid) setWeight()     { p.weight = p.a }

/*
baseA int
*/
func (p *Piramid) GetA() int { return p.a }

/*
baseA int
*/
func (p *Piramid) SetA(value int) {
	p.a = value
	p.setWeight()
}

/*
sideB int
*/
func (p *Piramid) GetB() int { return p.b }

/*
sideB int
*/
func (p *Piramid) SetB(value int) { p.b = value }

/*
height int
*/
func (p *Piramid) GetH() int { return p.h }

/*
height int
*/
func (p *Piramid) SetH(value int) { p.h = value }
