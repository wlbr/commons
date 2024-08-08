package stack

import (
	"fmt"
	"testing"
)

func prepareStack() *Stack[int] {
	q := NewStack[int]()
	q.Push(10)
	q.Push(20)
	q.Push(30)
	q.Push(40)
	q.Push(50)
	q.Push(60)
	q.Push(70)
	q.Push(80)
	q.Push(90)
	return q
}

func TestStackPush(t *testing.T) {
	s := prepareStack()

	s.Push(6)
	if s.len != 10 {
		t.Errorf("Length should be 10, but is %d", s.len)
	}
	if s.bottom.val != 10 {
		t.Errorf("First element should be 10, but is %d", s.bottom.val)
	}
	if s.top.val != 6 {
		t.Errorf("Last element should be 6, but is %d", s.top.val)
	}
}

func TestStackLength(t *testing.T) {
	s := prepareStack()
	if s.Len() != 9 {
		t.Errorf("Length should be 9, but is %d", s.Len())
	}
}

func TestStackPop(t *testing.T) {
	q := NewStack[int]()
	q.Push(10)
	q.Push(20)

	i, ok := q.Pop()
	if !ok {
		t.Errorf("Pop on populated stack unsuccessful. Stack length %d", q.len)
	}
	if i != 20 {
		t.Errorf("Pop expected value is 20, but got %d", i)
	}
	i, ok = q.Pop()
	if !ok {
		t.Errorf("Pop on populated stack unsuccessful. Stack length %d", q.len)
	}
	if i != 10 {
		t.Errorf("Pop expected value is 10, but got %d", i)
	}
	i, ok = q.Pop()
	if ok {
		t.Errorf("Pop on empty stack successful, error. Stack length %d", q.len)
	}

}

func TestStackPeek(t *testing.T) {
	s := NewStack[int]()
	s.Push(10)
	s.Push(20)

	i, ok := s.Peek()
	if i != 20 {
		t.Errorf("First peek should be 20, but is %d", i)
	}
	if !ok {
		t.Errorf("Peek on populated stack unsuccessful. Stack length %d", s.len)
	}

	i, ok = s.Peek()
	if !ok {
		t.Errorf("Peek on populated stack unsuccessful. Stack length %d", s.len)
	}
	if i != 20 {
		t.Errorf("Second peek should be 20 again, but is %d", i)
	}
	s.Pop()
	s.Pop()
	i, ok = s.Peek()
	fmt.Println(s)
	if ok {
		t.Errorf("Peek on empty stack successful. Stack length %d", s.len)
	}

}

func TestStackEmpty(t *testing.T) {
	s := NewStack[int]()
	if s.IsNotEmpty() {
		t.Errorf("New stack should be empty")
	}
	if !s.IsEmpty() {
		t.Errorf("New stack should be empty")
	}

	s = prepareStack()
	if !s.IsNotEmpty() {
		t.Errorf("Prepared stack should not be empty")
	}
	if s.IsEmpty() {
		t.Errorf("Prepared stack should not be empty")
	}

	s.Clear()
	if s.IsNotEmpty() {
		t.Errorf("Cleared stack should be empty")
	}
	if !s.IsEmpty() {
		t.Errorf("Cleared stack should be empty")
	}
}
