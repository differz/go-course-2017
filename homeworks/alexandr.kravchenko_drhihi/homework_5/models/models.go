package models

type Shaper interface {
	getWeight() int
}

type Shape []Shaper

func (s Shape) Len() int           { return len(s) }
func (s Shape) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Shape) Less(i, j int) bool { return s[i].getWeight() < s[j].getWeight() }
