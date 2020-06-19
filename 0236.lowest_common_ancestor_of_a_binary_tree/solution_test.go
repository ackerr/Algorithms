package leetcode0236

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	p := utils.TreeNode{Val: 1}
	left := utils.TreeNode{Val: 2, Left: &p, Right: &utils.TreeNode{Val: 3}}
	right := utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 7}}
	node := utils.TreeNode{Val: 4, Left: &left, Right: &right}
	if lowestCommonAncestor(&node, &p, &right) != &node {
		t.Errorf("result error")
	}
}
