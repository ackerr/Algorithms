package main

import "testing"

func TestMinPathSum(t *testing.T){
	grid := [][]int{{1,3,1}, {1,5,1}, {4,2,1}}
	result := minPathSum(grid)
	if result != 7{
		t.Errorf("results should be 7, not %d", result)
	}
}