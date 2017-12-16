package shapes

import (
	"fmt"
	"math"
	"reflect"
)

type Cone struct {
	m int
	r int
	h int
}

func NewCone(x int) (Shaper, error) {
	r := x / 2
	h := x
	v := math.Pi * float64(r*r*h) / 3
	m := int(K * v)
	return &Cone{m: m, r: r, h: h}, nil
}

func (s *Cone) Weight() int {
	return s.m
}

func (s *Cone) String() string {
	return fmt.Sprintf("(%s m:%d)", reflect.TypeOf(*s).Name(), s.Weight())
}
