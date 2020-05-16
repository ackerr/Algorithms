package leetcode0046

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	req := []int{1, 2}
	res := [][]int{{1, 2}, {2, 1}}
	if reflect.DeepEqual(permute(req), res) {
		t.Errorf("resutl should be %v", res)
	}

	req = []int{1, 2, 3}
	res = [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	if reflect.DeepEqual(permute(req), res) {
		t.Errorf("resutl should be %v", res)
	}
}
