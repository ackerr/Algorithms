package leetcode0230

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestKthSmallest(t *testing.T) {

	left := utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 3}}
	right := utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 7}}
	node := utils.TreeNode{Val: 4, Left: &left, Right: &right}
	if kthSmallest(&node, 4) != 4 {
		t.Errorf("4th smallest value should be 4")
	}
}
