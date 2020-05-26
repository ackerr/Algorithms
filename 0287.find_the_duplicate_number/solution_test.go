package leetcode0287

import "testing"

func TestFindDuplicate(t *testing.T) {
	req := [][]int{
		{1, 3, 3, 2, 4},
		{1, 2, 2, 3, 4},
		{2, 1, 1, 1, 1, 1},
	}

	rep := []int{3, 2, 1}
	for i, r := range req {
		if findDuplicate(r) != rep[i] {
			t.Errorf("%+v the duplicate num should be %d", r, rep[i])
		}
	}
}
