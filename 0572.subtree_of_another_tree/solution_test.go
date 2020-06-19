package leetcode0572

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestIsSubTree(t *testing.T) {

	p := utils.TreeNode{Val: 1}
	left := utils.TreeNode{Val: 2, Left: &p, Right: &utils.TreeNode{Val: 3}}
	right := utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 7}}
	node := utils.TreeNode{Val: 4, Left: &left, Right: &right}
	if isSubtree(&node, &left) == false {
		t.Errorf("results should be true")
	}

	node = utils.TreeNode{Val: 2, Left: &p, Right: &utils.TreeNode{Val: 3, Left: &p, Right: nil}}
	if isSubtree(&node, &left) == true {
		t.Errorf("results should be false")
	}
}
