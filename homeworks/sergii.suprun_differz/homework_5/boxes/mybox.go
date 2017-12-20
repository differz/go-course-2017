package boxes

import (
	"log"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/MastersAcademy/go-course-2017/homeworks/sergii.suprun_differz/homework_5/shapes"
)

type StateType int

const (
	EmptyState StateType = iota
	ByCorner1
	ByCorner2
	ByCorner3
	ByCorner4
	RandomState
)

type MyBox struct {
	n     int
	x     int
	state StateType
	box   []shapes.Shaper
	len   int
	q     int
}

func NewMyBox(n int, x int) (BlackBoxer, error) {
	return &MyBox{
			n:     n,
			x:     x,
			state: EmptyState,
			box:   []shapes.Shaper{},
			len:   n / x,
			q:     (n / x) * (n / x),
		},
		nil
}

func (b *MyBox) IsEmpty() bool {
	return len(b.box) == 0
}

func (b *MyBox) GetState() StateType {
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
	b.state = RandomState
	for i := 0; i < b.q; i++ {
		j := rand.Intn(len(names))
		shape, err := shapes.Create(names[j], b.x)
		if err != nil {
			log.Panicf("Cant create %s: ", names[j])
		}
		b.box = append(b.box, shape)
	}
}

func (b *MyBox) Shake(corner StateType) {
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
			case ByCorner1:
				mb[z+i] = dir
				mb[x+y] = rev
			case ByCorner2:
				mb[y+z] = dir
				mb[x+i] = rev
			case ByCorner3:
				mb[x+i] = dir
				mb[y+z] = rev
			case ByCorner4:
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
