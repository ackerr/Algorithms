package main

import "testing"

func TestMaxProduct(t *testing.T) {
	params := [][]int{
		{2, 4, 2, -4},
		{2, 0, -2, -4},
		{0, 2},
		{-3, -4, -5},
	}
	response := []int{16, 8, 2, 20}
	for i, param := range params {
		res := maxProduct(param)
		if res != response[i] {
			t.Errorf("params %v results should be %d, not %d", param, response[i], res)
		}
	}

}
