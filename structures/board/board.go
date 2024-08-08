package board

import (
	"fmt"

	"github.com/wlbr/commons/strategies/astar"
	"github.com/wlbr/commons/tools"
)

type element[T comparable] struct {
	val   T
	pos   []int
	board *BoardOfSquares[T]
}

func (e element[T]) String() string {
	return fmt.Sprintf("%v", e.val)
}

type BoardOfSquares[T comparable] struct {
	xdim, ydim int
	elems      [][]*element[T]
}

func (b BoardOfSquares[T]) String() string {
	s := ""
	for y := 0; y < b.ydim; y++ {
		for x := 0; x < b.xdim; x++ {
			s = s + fmt.Sprint(b.elems[y][x].val)
		}
		s = s + fmt.Sprintln()
	}
	return s
}

func NewBoardOFSquares[T comparable](xdim int, ydim int, def T) *BoardOfSquares[T] {
	b := &BoardOfSquares[T]{xdim: xdim, ydim: ydim}
	for y := 0; y < b.ydim; y++ {
		b.elems = append(b.elems, make([]*element[T], xdim, xdim))
	}
	for y, line := range b.elems {
		for x, _ := range line {
			b.elems[y][x] = &element[T]{pos: []int{x, y}, val: def, board: b}
		}
	}
	return b
}

func (b BoardOfSquares[T]) GetElement(x, y int) (*element[T], error) {
	if x < 0 || y < 0 || x >= b.xdim || y >= b.ydim {
		return nil, fmt.Errorf("position [%d,%d] is out of bounds", x, y)
	}
	return b.elems[y][x], nil
}

func (b BoardOfSquares[T]) Get(x, y int) (T, error) {
	if x < 0 || y < 0 || x >= b.xdim || y >= b.ydim {
		return b.elems[0][0].val, fmt.Errorf("position [%d,%d] is out of bounds", x, y)
	}
	return b.elems[y][x].val, nil
}

func (b BoardOfSquares[T]) Set(x, y int, v T) error {
	if x < 0 || y < 0 || x >= b.xdim || y > b.ydim {
		return fmt.Errorf("position [%d,%d] is out of bounds", x, y)
	}
	b.elems[y][x].val = v
	return nil
}

func getDirections() (directions [][]int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				directions = append(directions, []int{j, i})
			}
		}
	}
	return directions
}

func (b BoardOfSquares[T]) GetNeighborCoordinates(x, y int) (neighbors [][]int, e error) {
	if x < 0 || y < 0 || x >= b.xdim || y > b.ydim {
		return nil, fmt.Errorf("position [%d,%d] is out of bounds.", x, y)
	}
	directions := getDirections()
	for _, d := range directions {
		if x+d[0] >= 0 && x+d[0] < b.xdim && y+d[1] >= 0 && y+d[1] < b.ydim {
			n := []int{x + d[0], y + d[1]}
			neighbors = append(neighbors, n)
		}
	}
	return neighbors, nil
}

func (b BoardOfSquares[T]) GetNeighborElements(x, y int) (neighbors []astar.Pather, e error) {
	coords, e := b.GetNeighborCoordinates(x, y)
	if e != nil {
		return nil, e
	}
	for _, coordinate := range coords {
		elem, _ := b.GetElement(coordinate[0], coordinate[1])
		neighbors = append(neighbors, elem)
	}

	return neighbors, nil
}

func (r *element[T]) ManhattanDistance(to *element[T]) float64 {
	return float64(tools.Abs(r.pos[0]-to.pos[0]) + tools.Abs(r.pos[1]-to.pos[1]))
}

func (r *element[T]) PathNeighbors() []astar.Pather {
	n, _ := r.board.GetNeighborElements(r.pos[0], r.pos[1])
	return n
}

func (from *element[T]) PathNeighborCost(to astar.Pather) float64 {
	return float64(1)
}

func (from *element[T]) PathEstimatedCost(to astar.Pather) float64 {
	return from.ManhattanDistance(to.(*element[T]))
}
