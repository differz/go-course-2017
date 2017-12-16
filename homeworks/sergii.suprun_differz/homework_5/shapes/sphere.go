package shapes

import (
	"fmt"
	"math"
	"reflect"
)

type Sphere struct {
	m int
	r int
}

func NewSphere(x int) (Shaper, error) {
	r := x / 2
	v := float64(r*r*r) * math.Pi * 4 / 3
	m := int(K * v)
	return &Sphere{m: m, r: r}, nil
}

func (s *Sphere) Weight() int {
	return s.m
}

func (s *Sphere) String() string {
	return fmt.Sprintf("(%s m:%d)", reflect.TypeOf(*s).Name(), s.Weight())
}
