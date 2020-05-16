package leetcode0033

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{6, 7, 8, 9, 1, 3, 4, 5}
	target := 8
	index := search(nums, target)
	if index != 2 {
		t.Errorf("target %d index error", target)
	}

	nums = []int{16, 27, 38, 49, 1, 3, 4, 5, 10, 13}
	target = 8
	index = search(nums, target)
	if index != -1 {
		t.Errorf("target %d index error", target)
	}
}
