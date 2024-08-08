package torus

import (
	"fmt"
	"log"
)

type Torus struct {
	data             [][]interface{}
	dimX, dimY       int
	defaultValue     interface{}
	elementFormatter func(interface{}) string
	comparator       func(interface{}, interface{}) bool
}

func NewTorus(dimX, dimY int, defaultValue interface{}) *Torus {
	a := &Torus{
		data:         make([][]interface{}, dimY),
		dimX:         dimX,
		dimY:         dimY,
		defaultValue: interface{}(defaultValue),
	}
	for i := range a.data {
		a.data[i] = make([]interface{}, dimX)
		for j := range a.data[i] {
			a.data[i][j] = a.defaultValue
		}
	}

	return a
}

func NewTorusFromString(s []string, defaultValue interface{}) *Torus {
	if len(s) > 0 {
		a := &Torus{}
		a.dimY = len(s)
		a.dimX = len(s[0])
		a.defaultValue = defaultValue
		a.elementFormatter = func(e interface{}) string {
			return fmt.Sprintf("%c", e)
		}
		a.comparator = func(a, b interface{}) bool {
			return a.(rune) == b.(rune)
		}

		for _, line := range s {
			l := make([]interface{}, len(line))
			a.data = append(a.data, l)
			for i, c := range line {
				l[i] = c
			}
		}
		return a
	}
	log.Fatal("NewTorusFromString: cannot create from empty string")
	return nil
}

func (t *Torus) String() string {
	s := ""
	for _, line := range t.data {
		for _, c := range line {
			s += t.elementFormatter(c)
		}
		s += "\n"
	}
	return s
}

func (t *Torus) ShallowEquals(at *Torus) bool {
	if at.dimX != t.dimX || at.dimY != t.dimY {
		return false
	}
	for y := 0; y < t.dimY; y++ {
		for x := 0; x < t.dimX; x++ {
			if at.Get(x, y) != t.Get(x, y) {
				return false
			}
		}
	}
	return true
}

func (t *Torus) Equals(at *Torus) bool {
	if at.dimX != t.dimX || at.dimY != t.dimY {
		return false
	}
	for y := 0; y < t.dimY; y++ {
		for x := 0; x < t.dimX; x++ {
			if !t.comparator(at.Get(x, y), t.Get(x, y)) {
				return false
			}
		}
	}
	return true
}

func (t *Torus) Copy() *Torus {
	a := NewTorus(t.dimX, t.dimY, t.defaultValue)
	a.elementFormatter = t.elementFormatter
	for y := 0; y < t.dimY; y++ {
		for x := 0; x < t.dimX; x++ {
			a.Set(x, y, t.Get(x, y))
		}
	}
	return a
}

func (t *Torus) Set(x, y int, value interface{}) {
	t.data[y%t.dimY][x%t.dimX] = value
}

func (t *Torus) Get(x, y int) interface{} {
	return t.data[y%t.dimY][x%t.dimX]
}

func (t *Torus) GetInt(x, y int) int {
	return t.Get(x, y).(int)
}

func (t *Torus) GetBool(x, y int) bool {
	return t.Get(x, y).(bool)
}

func (t *Torus) GetString(x, y int) string {
	return fmt.Sprintf(t.elementFormatter(t.Get(x, y)))
}

func (t *Torus) GetRune(x, y int) rune {
	return t.Get(x, y).(rune)
}

func (t *Torus) GetByte(x, y int) byte {
	return t.Get(x, y).(byte)
}
