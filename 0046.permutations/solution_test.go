package leetcode0046

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	req := []int{1, 2}
	res := [][]int{{2, 1}, {1, 2}}
	if !reflect.DeepEqual(permute(req), res) {
		t.Errorf("resutl should be %v", res)
	}

	req = []int{1, 2, 3}
	res = [][]int{{3, 2, 1}, {2, 3, 1}, {3, 1, 2}, {1, 3, 2}, {2, 1, 3}, {1, 2, 3}}
	if !reflect.DeepEqual(permute(req), res) {
		t.Errorf("resutl should be %v", res)
	}
}
