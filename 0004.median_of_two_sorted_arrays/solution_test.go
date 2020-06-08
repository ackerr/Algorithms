package leetcode0004

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	req1 := [][]int{
		{1, 2, 3, 4},
		{1, 2, 4, 5},
		{1},
		{},
	}
	req2 := [][]int{
		{5, 6, 7, 8},
		{1, 2, 3, 4},
		{2, 3},
		{1, 2},
	}

	res := []float64{4.5, 2.5, 2.0, 1.5}
	for i, r := range req1 {
		if findMedianSortedArrays(r, req2[i]) != res[i] {
			t.Errorf("findMedianSortedArrays(%+v, %+v) should be %.1f", r, req2[i], res[i])
		}
	}

}
