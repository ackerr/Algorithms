package leetcode0096

import (
	"testing"
)

func TestNumTree(t *testing.T) {
	num := numTrees(3)
	if num != 5 {
		t.Errorf("results error")
	}

	num = numTrees(5)
	if num != 42 {
		t.Errorf("results error")
	}

}
