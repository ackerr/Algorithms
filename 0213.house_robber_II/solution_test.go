package main

import "testing"

func TestRop(t *testing.T) {
	params := []int{1,2,3,6,8,5,5,8,3,1}
	r := rob(params)
	if r != 22{
		t.Errorf("results should be %d", r)
	}
}
