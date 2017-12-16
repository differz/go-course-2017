package shapes

import (
	"fmt"
	"reflect"
)

type Cube struct {
	m int
	a int
}

func NewCube(x int) (Shaper, error) {
	a := x
	v := float64(a * a * a)
	m := int(K * v)
	return &Cube{m: m, a: a}, nil
}

func (s *Cube) Weight() int {
	return s.m
}

func (s *Cube) String() string {
	return fmt.Sprintf("(%s m:%d)", reflect.TypeOf(*s).Name(), s.Weight())
}
