package list

// List is the common interface implemented by LinkedList
type List interface {
	Add(e interface{})
	Insert(index int, e interface{})
<<<<<<< HEAD
	Set(index int, e interface{}) error
	Remove(e interface{}) bool
	Get(index int) (interface{}, error)
	Size() int
	IsEmpty() bool
	String() string
=======
	Set(index int, e interface{})
	Remove(e interface{}) bool
	Get(index int) interface{}
	Size() int
	IsEmpty() int
>>>>>>> MastersAcademy/master
}
