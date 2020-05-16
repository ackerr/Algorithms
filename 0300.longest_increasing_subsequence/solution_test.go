package leetcode0300

import "testing"

func TestLengthOfLIS(t *testing.T) {
	req := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{1, 9, 2, 5, 3, 7, 101, 18},
		{1, 9, 2, 5, 3, 7, 18, 101},
		{1, 3, 2, 5, 3, 2, 11, 18},
	}
	res := []int{4, 5, 6, 5}
	for i, r := range req {
		result := lengthOfLIS(r)
		if result != res[i] {
			t.Errorf("LengthOfLIS(%v), result should be %d, not %d", req[i], res[i], result)
		}
	}

}
