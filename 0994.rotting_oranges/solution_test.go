package main

import "testing"

func TestOrangesRotting(t *testing.T) {
	r := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	if orangesRotting(r) != 4 {
		t.Errorf("result should be 4")
	}

	r = [][]int{{2, 1, 1}, {1, 1, 0}, {0, 0, 1}}
	if orangesRotting(r) != -1 {
		t.Errorf("result should be -1")
	}

}
