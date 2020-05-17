package leetcode0230

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	left := utils.TreeNode{2, &utils.TreeNode{1, nil, nil}, &utils.TreeNode{3, nil, nil}}
	right := utils.TreeNode{6, &utils.TreeNode{5, nil, nil}, &utils.TreeNode{7, nil, nil}}
	node := utils.TreeNode{4, &left, &right}
	if kthSmallest(&node, 4) != 4 {
		t.Errorf("4th smallest value should be 4")
	}
}
