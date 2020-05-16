package leetcode0016

import "testing"

func TestThreeSumClosest(t *testing.T) {
	req := []int{1, 4, 5, 6, 8, 2, -5, 2, 5}
	if threeSumClosest(req, 0) != 0 {
		t.Errorf("result should be 0")
	}

}
