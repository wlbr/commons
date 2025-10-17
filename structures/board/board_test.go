package board

import (
	"fmt"
	"testing"
)

func (b BoardOfSquares[T]) checkErrValues(x, y int, v T) bool {
	r, e := b.Get(x, y)
	return (e != nil) || (r != v)
}

func TestBoardValues(t *testing.T) {

	b := NewBoardOFSquares[string](10, 3, "-")

	b.Set(2, 0, "3-1")
	b.Set(2, 1, "3-2")
	b.Set(4, 2, "5-3")

	if b.checkErrValues(2, 0, "3-1") {
		t.Fail()
	}
	if b.checkErrValues(2, 1, "3-2") {
		t.Fail()
	}
	if b.checkErrValues(2, 2, "-") {
		t.Fail()
	}
	if b.checkErrValues(4, 2, "5-3") {
		t.Fail()
	}
}

func (b BoardOfSquares[T]) printNeighbors(x, y int) {
	n, _ := b.GetNeighborCoordinates(x, y)
	fmt.Printf("[%d,%d] - %v\n", x, y, n)
}

func (b BoardOfSquares[T]) checkErrNeighbors(x, y int, expected int) bool {
	r, e := b.GetNeighborCoordinates(x, y)
	if e != nil {
		return false
	}
	return len(r) == expected
}

func TestBoardNeighbors(t *testing.T) {

	b := NewBoardOFSquares[string](10, 3, "-")

	b.Set(2, 0, "3-1")
	b.Set(2, 1, "3-2")
	b.Set(4, 2, "5-3")

	//center
	b.checkErrNeighbors(1, 1, 8)
	//borders
	b.checkErrNeighbors(1, 0, 5)
	b.checkErrNeighbors(9, 1, 5)
	b.checkErrNeighbors(1, 2, 5)
	b.checkErrNeighbors(0, 1, 5)

	//corners
	b.checkErrNeighbors(0, 0, 3)
	b.checkErrNeighbors(9, 0, 3)
	b.checkErrNeighbors(9, 2, 3)
	b.checkErrNeighbors(0, 2, 3)

}
