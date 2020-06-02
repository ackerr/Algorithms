package leetcode0121

import "testing"

func TestMaxProfit(t *testing.T) {
	req := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 1, 5, 3, 6, 4},
		{7, 6, 5, 4, 3, 2},
	}

	res := []int{5, 5, 0}
	for i, r := range req {
		if maxProfit(r) != res[i] {
			t.Errorf("maxProfit(%+v) should be %d, not %d", r, res[i], maxProfit(r))

		}
	}

}
