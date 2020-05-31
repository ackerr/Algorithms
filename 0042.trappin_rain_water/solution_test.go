package leetcode0042

import "testing"

func TestTrap(t *testing.T) {
	req := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{0, 2, 0},
		{2, 0, 2},
	}

	res := []int{6, 0, 2}

	for i, r := range req {
		if trap(r) != res[i] {
			t.Errorf("trap(%+v) 's results should be %d, not %d", r, res[i], trap(r))
		}
	}

}
