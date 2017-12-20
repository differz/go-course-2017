package boxes

import (
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/MastersAcademy/go-course-2017/homeworks/sergii.suprun_differz/homework_5/shapes"
)

const (
	EMPTY = iota
	CORNER1
	CORNER2
	CORNER3
	CORNER4
	RANDOM
)

type MyBox struct {
	n     int
	x     int
	state int
	box   []shapes.Shaper
	len   int
	q     int
}

func NewMyBox(n int, x int) (BlackBoxer, error) {
	return &MyBox{
			n:     n,
			x:     x,
			state: EMPTY,
			box:   []shapes.Shaper{},
			len:   n / x,
			q:     (n / x) * (n / x),
		},
		nil
}

func (b *MyBox) IsEmpty() bool {
	return len(b.box) == 0
}

func (b *MyBox) GetState() int {
	return b.state
}

func (b *MyBox) PrintWeight() string {
	k := b.n / b.x
	s := ""
	for key, val := range b.box {
		s += strconv.Itoa(val.Weight()) + " "
		if (key+1)%k == 0 {
			s += "\n"
		}
	}
	return s
}

func (b *MyBox) String() string {
	return "(box x=" + strconv.Itoa(b.x) + ")"
}

func (b *MyBox) Generate() {
	names := shapes.GetAvailable()
	rand.Seed(time.Now().UnixNano())
	b.state = RANDOM
	for i := 0; i < b.q; i++ {
		j := rand.Intn(len(names))
		shape, _ := shapes.Create(names[j], b.x)
		b.box = append(b.box, shape)
	}
}

func (b *MyBox) Shake(corner int) {
	index := 0
	mb := make([]shapes.Shaper, b.q)
	sort.Sort(shapes.ByWeight(b.box))

	for n := 0; n < b.len; n++ {
		for i, j := 0, n; i <= n; i, j = i+1, j-1 {
			dir := b.box[index]
			rev := b.box[b.q-index-1]
			index++

			x := (b.len - 1 - j) * b.len
			y := b.len - 1 - i
			z := j * b.len

			switch corner {
			case CORNER1:
				mb[z+i] = dir
				mb[x+y] = rev
			case CORNER2:
				mb[y+z] = dir
				mb[x+i] = rev
			case CORNER3:
				mb[x+i] = dir
				mb[y+z] = rev
			case CORNER4:
				mb[x+y] = dir
				mb[z+i] = rev
			default:
				return
			}
		}
	}
	b.state = corner
	b.box = mb
}
