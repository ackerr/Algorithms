package leetcode0011

import "testing"

func TestMaxArea(t *testing.T) {
	req := [][]int{
		{1, 1, 1, 1, 5, 5, 1},
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
	}

	r := []int{6, 49}

	for i, v := range req {
		res := maxArea(v)
		if res != r[i] {
			t.Errorf("result shoude be %d, not %d", res, r[i])
		}
	}
}
