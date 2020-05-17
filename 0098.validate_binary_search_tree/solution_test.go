package leetcode0098

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestIsValidBST(t *testing.T) {

	p := utils.TreeNode{1, nil, nil}
	left := utils.TreeNode{2, &p, &utils.TreeNode{3, nil, nil}}
	right := utils.TreeNode{6, &utils.TreeNode{5, nil, nil}, &utils.TreeNode{7, nil, nil}}
	node := utils.TreeNode{4, &left, &right}

	if !isValidBST(&node) {
		t.Errorf("%+v is a BST", node)
	}

	p.Val = 2
	if isValidBST(&node) {
		t.Errorf("%+v is not a BST", node)
	}

}
