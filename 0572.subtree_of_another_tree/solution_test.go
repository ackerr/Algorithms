package leetcode0572

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestIsSubTree(t *testing.T) {
	p := utils.TreeNode{1, nil, nil}
	left := utils.TreeNode{2, &p, &utils.TreeNode{3, nil, nil}}
	right := utils.TreeNode{6, &utils.TreeNode{5, nil, nil}, &utils.TreeNode{7, nil, nil}}
	node := utils.TreeNode{4, &left, &right}
	if isSubtree(&node, &left) == false {
		t.Errorf("results should be true")
	}

	node = utils.TreeNode{2, &p, &utils.TreeNode{3, &p, nil}}
	if isSubtree(&node, &left) == true {
		t.Errorf("results should be false")
	}
}
