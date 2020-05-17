package leetcode0100

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	p := utils.TreeNode{1, nil, nil}
	left := utils.TreeNode{2, &p, &utils.TreeNode{3, nil, nil}}
	right := utils.TreeNode{6, &utils.TreeNode{5, nil, nil}, &utils.TreeNode{7, nil, nil}}
	if isSameTree(&right, &right) == false {
		t.Errorf("results should be true")
	}

	if isSameTree(&right, &left) == true {
		t.Errorf("results should be false")
	}
}
