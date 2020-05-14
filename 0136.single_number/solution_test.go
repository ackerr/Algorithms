package main

import "testing"

func TestSingleNumber(t *testing.T) {
	req := [][]int{
		{1, 1, 2, 3, 3},
		{1, 1, 3, 3, 4},
		{01, 2, 3, 2, 0, 1},
	}
	res := []int{2, 4, 3}
	for i, r := range req {
		if singleNumber(r) != res[i] {
			t.Errorf("nums is %+v, result shoule be %+v ", r, res[i])
		}
	}
}
