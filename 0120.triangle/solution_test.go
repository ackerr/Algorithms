package leetcode0120

import "testing"

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	if minimumTotal(triangle) != 11 {
		t.Errorf("results should be 11")
	}
}
