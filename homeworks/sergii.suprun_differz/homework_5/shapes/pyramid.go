package shapes

import (
	"fmt"
	"reflect"
)

type Pyramid struct {
	m int
	a int
	h int
}

func NewPyramid(x int) (Shaper, error) {
	a := x
	h := x
	v := float64(a*a*h) / 3
	m := int(K * v)
	return &Pyramid{m: m, a: a, h: h}, nil
}

func (s *Pyramid) Weight() int {
	return s.m
}

func (s *Pyramid) String() string {
	return fmt.Sprintf("(%s m:%d)", reflect.TypeOf(*s).Name(), s.Weight())
}
