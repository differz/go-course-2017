package LinkedList

import (
	"errors"
)

type LinkedList struct {
	value interface{}
	prev  *LinkedList
	next  *LinkedList
}

func (l *LinkedList) getData(index int) (rl *LinkedList, i int, v interface{}, err error) {

	for cl := l; cl != nil && (index == -1 || i <= index); cl, i = cl.next, i+1 {
		v = cl.value
		rl = cl
	}

	if index != -1 && i-1 < index {
		err = errors.New("index out of range")
	}

	return
}

func Len(l *LinkedList) (i int) {
	_, i, _, _ = l.getData(-1)
	return
}

func (l *LinkedList) Get(index uint) (v interface{}, err error) {
	_, _, v, err = l.getData(int(index))
	if err != nil {
		v = nil
	}
	return
}

func (l *LinkedList) Set(index uint, value interface{}) (err error) {
	cl, _, _, err := l.getData(int(index))
	if err == nil {
		cl.value = value
	}
	return
}

func Append(l *LinkedList, value interface{}) *LinkedList {
	if l == nil {
		return &LinkedList{value: value}
	}
	cl, _, _, _ := l.getData(-1)
	nl := LinkedList{value: value, prev: cl}
	cl.next = &nl
	return l
}

func (l *LinkedList) Delete(index uint) (err error){
	cl, _, _, err := l.getData(int(index))
	switch {
	default:
		prev, next := cl.prev, cl.next
		prev.next = next
		if next != nil {
			next.prev = prev
		}
		cl.next, cl.prev = nil, nil
	case err != nil:
	case l == cl:
		n := l.next
		l.value = n.value
		l.next = n.next
		l.prev = nil
		n.prev, n.next = nil, nil
	}
	return
}

func (l *LinkedList) Value() (res []interface{}) {
	for cl := l; cl != nil; cl = cl.next {
		res = append(res, cl.value)
	}
	return
}

func New(value interface{}) *LinkedList {
	return &LinkedList{value:value}
}
