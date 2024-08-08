package board

import (
	"fmt"
	"log"
	"testing"

	"github.com/wlbr/commons/strategies/astar"
)

func (b BoardOfSquares[T]) printNeighbors(x, y int) {
	n, _ := b.GetNeighborCoordinates(x, y)
	fmt.Printf("[%d,%d] - %v\n", x, y, n)
}

func (b BoardOfSquares[T]) checkErr(x, y int, v T) bool {
	r, e := b.Get(x, y)
	if e != nil {
		return false
	}
	return r == v
}

func TestBoard(t *testing.T) {

	b := NewBoardOFSquares[string](10, 3, "-")

	b.Set(2, 0, "3-1")
	b.Set(2, 1, "3-2")
	b.Set(4, 2, "5-3")

	if b.checkErr(2, 0, "3-1") {
		t.Fail()
	}
	if b.checkErr(2, 1, "3-2") {
		t.Fail()
	}
	if b.checkErr(2, 0, "3-2") {
		t.Fail()
	}
	if b.checkErr(4, 2, "5-3") {
		t.Fail()
	}

	fmt.Println(b)

	b.printNeighbors(1, 1)
	b.printNeighbors(0, 1)
	b.printNeighbors(0, 2)

	b.printNeighbors(8, 1)
	b.printNeighbors(9, 1)
	b.printNeighbors(9, 2)

	from, _ := b.GetElement(0, 0)
	to, _ := b.GetElement(b.xdim-1, b.ydim-1)
	p, distance, found := astar.Path(from, to)

	if !found {
		log.Println("Could not find path")
	}
	for _, v := range p {
		e := v.(*element[string])
		fmt.Println(e.pos)
	}
	fmt.Println(distance)

}
