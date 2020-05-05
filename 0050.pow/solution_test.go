package main

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	req := [][]float64{{19, 11}, {19, -11}, {999, 10}, {999, 11}, {-10, 0}}

	for _, r := range req {
		value := math.Pow(r[0], r[1])
		if myPow(r[0], int(r[1])) != value {
			t.Errorf("pow(%f, %f) should be %f", r[0], r[1], value)
		}
	}
}
