package stack

import "fmt"

type element[T comparable] struct {
	val   T
	next  *element[T]
	last  *element[T]
	Stack *Stack[T]
}

func (e element[T]) String() string {
	return fmt.Sprintf("%v", e.val)
}

type Stack[T comparable] struct {
	len    int
	bottom *element[T]
	top    *element[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (q *Stack[T]) String() string {
	s := "["
	for e := q.bottom; e != nil; e = e.next {
		if e == q.bottom {
			s = fmt.Sprintf("%s%s", s, e)
		} else {
			s = fmt.Sprintf("%s,%s", s, e)
		}
	}
	return s + "]"
}

func (q *Stack[T]) Len() int {
	return q.len
}

func (q *Stack[T]) Push(val T) {
	n := &element[T]{val: val, Stack: q}
	if q.len == 0 {
		q.bottom = n
		q.top = n
	} else {
		q.top.next = n
		n.last = q.top
		q.top = n
	}
	q.len++
}

func (q *Stack[T]) Peek() (val T, ok bool) {
	if q.len == 0 {
		return val, false
	}
	return q.top.val, true
}

func (q *Stack[T]) Pop() (val T, ok bool) {
	if q.len == 0 {
		return val, false
	}
	val = q.top.val
	if q.len == 1 {
		q.bottom = nil
		q.top = nil
	} else {
		q.top = q.top.last
		q.top.next = nil
	}
	q.len--
	return val, true
}

func (q *Stack[T]) Clear() {
	q.bottom = nil
	q.top = nil
	q.len = 0
}

func (q *Stack[T]) IsEmpty() bool {
	return q.len == 0
}

func (q *Stack[T]) IsNotEmpty() bool {
	return q.len != 0
}
