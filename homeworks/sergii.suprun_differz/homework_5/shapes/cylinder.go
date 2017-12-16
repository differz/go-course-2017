package shapes

import (
	"fmt"
	"math"
	"reflect"
)

type Cylinder struct {
	m int
	r int
	h int
}

func NewCylinder(x int) (Shaper, error) {
	r := x / 2
	h := x
	v := math.Pi * float64(r*r*h)
	m := int(K * v)
	return &Cylinder{m: m, r: r, h: h}, nil
}

func (s *Cylinder) Weight() int {
	return s.m
}

func (s *Cylinder) String() string {
	return fmt.Sprintf("(%s m:%d)", reflect.TypeOf(*s).Name(), s.Weight())
}
