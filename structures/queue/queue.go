package queue

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type element[T constraints.Ordered] struct {
	val   T
	next  *element[T]
	last  *element[T]
	queue *Queue[T]
}

func (e element[T]) String() string {
	return fmt.Sprintf("%v", e.val)
}

type Queue[T constraints.Ordered] struct {
	len   int
	first *element[T]
	last  *element[T]
}

func NewQueue[T constraints.Ordered]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) String() string {
	s := "["
	for e := q.first; e != nil; e = e.next {
		if e == q.first {
			s = fmt.Sprintf("%s%s", s, e)
		} else {
			s = fmt.Sprintf("%s,%s", s, e)
		}
	}
	return s + "]"
}

func (q *Queue[T]) Len() int {
	return q.len
}

func (q *Queue[T]) First() T {
	return q.first.val
}

func (q *Queue[T]) Last() T {
	return q.last.val
}

func (q *Queue[T]) Get(i int) (T, error) {
	if i < 0 || i >= q.len {
		return q.first.val, fmt.Errorf("Index out of range")
	}
	e := q.first
	for j := 0; j < i; j++ {
		e = e.next
	}
	return e.val, nil
}

func (q *Queue[T]) Add(val T) {
	n := &element[T]{val: val, queue: q}
	if q.len == 0 {
		q.first = n
		q.last = n
	} else {
		q.last.next = n
		n.last = q.last
		q.last = n
	}
	q.len++
}

func (q *Queue[T]) Push(val T) {
	n := &element[T]{val: val, queue: q}
	if q.len == 0 {
		q.first = n
		q.last = n
	} else {
		q.first.last = n
		n.next = q.first
		q.first = n
	}
	q.len++
}

func (q *Queue[T]) RemoveFirst() {
	if q.len == 0 {
		return
	}
	if q.len == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.first = q.first.next
		q.first.last = nil
	}
	q.len--
}

func (q *Queue[T]) RemoveLast() {
	if q.len == 0 {
		return
	}
	if q.len == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.last = q.last.last
		q.last.next = nil
	}
	q.len--
}

func (q *Queue[T]) Remove(e *element[T]) {
	if e.queue != q {
		return
	}
	if e.last == nil {
		q.RemoveFirst()
	} else if e.next == nil {
		q.RemoveLast()
	} else {
		e.last.next = e.next
		e.next.last = e.last
		q.len--
	}
}

func (q *Queue[T]) RemoveFirstVal(val T) {
	for e := q.first; e != nil; e = e.next {
		if e.val == val {
			q.Remove(e)
			return
		}
	}
}

func (q *Queue[T]) RemoveAllVal(val T) {
	for e := q.first; e != nil; e = e.next {
		if e.val == val {
			q.Remove(e)
		}
	}
}

func (q *Queue[T]) Contains(val T) bool {
	for e := q.first; e != nil; e = e.next {
		if e.val == val {
			return true
		}
	}
	return false
}

func (q *Queue[T]) Position(val T) (pos int) {
	for e := q.first; e != nil; e = e.next {
		if e.val == val {
			return pos
		}
		pos++
	}
	return -1
}

func (q *Queue[T]) Clear() {
	q.first = nil
	q.last = nil
	q.len = 0
}

func (q *Queue[T]) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue[T]) IsNotEmpty() bool {
	return q.len != 0
}
