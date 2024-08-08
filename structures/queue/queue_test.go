package queue

import (
	"testing"
)

func prepareQueue() *Queue[int] {
	q := NewQueue[int]()
	q.Add(10)
	q.Add(20)
	q.Add(30)
	q.Add(40)
	q.Add(50)
	q.Add(60)
	q.Add(70)
	q.Add(80)
	q.Add(90)
	return q
}

func TestQueueAdd(t *testing.T) {
	q := prepareQueue()

	q.Push(6)
	if q.len != 10 {
		t.Errorf("Length should be 10, but is %d", q.len)
	}
	if q.first.val != 6 {
		t.Errorf("First element should be 6, but is %d", q.first.val)
	}
	if q.last.val != 90 {
		t.Errorf("Last element should be 90, but is %d", q.last.val)
	}
}

func TestQueueLength(t *testing.T) {
	q := prepareQueue()

	q.Push(6)
	if q.Len() != 10 {
		t.Errorf("Length should be 10, but is %d", q.len)
	}
}

func TestQueueRemove(t *testing.T) {
	q := prepareQueue()

	q.Push(6)

	q.RemoveLast()
	q.RemoveFirst()
	q.Remove(q.first.next)
	q.RemoveFirstVal(70)

	if q.Len() != 6 {
		t.Errorf("Length should be 6, but is %d", q.len)
	}
}

func TestQueueFirstVal(t *testing.T) {
	q := prepareQueue()

	q.Push(60)
	q.Push(70)
	q.Push(80)
	q.Add(60)

	q.RemoveAllVal(60)

	if q.Len() != 10 {
		t.Errorf("Length should be 10, but is %d", q.len)
	}
}

func TestQueueFind(t *testing.T) {
	q := prepareQueue()

	if !q.Contains(30) {
		t.Errorf("Queue should contain 30")
	}
	if q.Contains(35) {
		t.Errorf("Queue should not contain 35")
	}

	i := q.Position(30)
	if i != 2 {
		t.Errorf("Position of 30 should be 2, but is %d", i)
	}
	i = q.Position(35)
	if i != -1 {
		t.Errorf("Position of 35 should be -1 (not found), but is %d", i)
	}
}

func TestQueueGet(t *testing.T) {
	q := prepareQueue()

	i, e := q.Get(2)
	if i != 30 || e != nil {
		t.Errorf("Value at position 2 should be 30, but is %d", i)
	}

	i, e = q.Get(20)
	if e == nil {
		t.Errorf("No error though index %d greater than length %d", 20, q.Len())
	}
}

func TestQueueEmpty(t *testing.T) {
	q := NewQueue[int]()
	if q.IsNotEmpty() {
		t.Errorf("New queue should be empty")
	}
	if !q.IsEmpty() {
		t.Errorf("New queue should be empty")
	}

	q = prepareQueue()
	if !q.IsNotEmpty() {
		t.Errorf("Prepared queue should be empty")
	}
	if q.IsEmpty() {
		t.Errorf("Prepared queue should be empty")
	}

	q.Clear()
	if q.IsNotEmpty() {
		t.Errorf("New queue should be empty")
	}
	if !q.IsEmpty() {
		t.Errorf("New queue should be empty")
	}
}
