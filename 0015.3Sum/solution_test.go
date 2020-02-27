package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	req := []int{-3, 0, 3, 1, 0, 2}
	res := [][]int{{-3, 0, 3}, {-3, 1, 2}}
	r := reflect.DeepEqual(res, threeSum(req))
	if !r {
		t.Errorf("result shoude be [{-3, 0, 3}, {-3, 1, 2}]")
	}
}
