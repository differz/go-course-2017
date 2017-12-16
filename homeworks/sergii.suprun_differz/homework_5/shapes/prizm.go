package shapes

import (
	"fmt"
	"math"
	"reflect"
)

type Prizm struct {
	m int
	a int
	h int
}

func NewPrizm(x int) (Shaper, error) {
	a := x
	h := x
	v := float64(a*a*h) * math.Sqrt(3) / 4
	m := int(K * v)
	return &Prizm{m: m, a: a, h: h}, nil
}

func (s *Prizm) Weight() int {
	return s.m
}

func (s *Prizm) String() string {
	return fmt.Sprintf("(%s m:%d)", reflect.TypeOf(*s).Name(), s.Weight())
}
