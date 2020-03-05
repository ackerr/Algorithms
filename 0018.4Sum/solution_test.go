package main

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	req := []int{1, 0, -1, 0, -2, 2}
	res := [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	if !reflect.DeepEqual(fourSum(req, 0), res) {
		t.Errorf("result error")
	}
}
