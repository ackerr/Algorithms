package leetcode0098

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestIsValidBST(t *testing.T) {

	p := utils.TreeNode{Val: 1}
	left := utils.TreeNode{Val: 2, Left: &p, Right: &utils.TreeNode{Val: 3}}
	right := utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 7}}
	node := utils.TreeNode{Val: 4, Left: &left, Right: &right}

	if !isValidBST(&node) {
		t.Errorf("%+v is a BST", node)
	}

	p.Val = 2
	if isValidBST(&node) {
		t.Errorf("%+v is not a BST", node)
	}

}
