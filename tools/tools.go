package tools

import "github.com/wlbr/commons/log"

// CheckErr is a convenience function makes error handling dangerously simple.
func CheckErr(err error) {
	if err != nil {
		log.Debug("%s", err)
	}
}

// Minf64 returns the minimum of a slice of float64
func Minf64(v []float64) float64 {
	m := v[0]
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}

// Maxf64 returns the maximum of a slice of float64
func Maxf64(v []float64) float64 {
	m := v[0]
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
