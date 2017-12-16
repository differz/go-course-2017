package shapes

// K - density: m = ρV
const K = 1.0

// Shape is common interface for all shapes
type Shaper interface {
	Weight() int
	String() string
}

// ByWeight implements sort.Interface for []Shaper based on the Weight()
type ByWeight []Shaper

func (a ByWeight) Len() int {
	return len(a)
}
func (a ByWeight) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByWeight) Less(i, j int) bool {
	return a[i].Weight() < a[j].Weight()
}
