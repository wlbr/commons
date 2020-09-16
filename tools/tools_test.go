package tools

import (
	"testing"
)

func TestMinf64(t *testing.T) {
	data := []float64{3.1, 3.5, 2.1, 22.2, 32, 1, -2.3}
	if Minf64(data) != -2.3 {
		t.Fail()
	}
}

func TestMaxf64(t *testing.T) {
	data := []float64{3.1, 3.5, 2.1, 22.2, 32, 1, -2.3}
	if Maxf64(data) != 32 {
		t.Fail()
	}
}
