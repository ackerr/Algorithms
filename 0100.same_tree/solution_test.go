package leetcode0100

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestIsSameTree(t *testing.T) {

	p := utils.TreeNode{Val: 1}
	left := utils.TreeNode{Val: 2, Left: &p, Right: &utils.TreeNode{Val: 3}}
	right := utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 7}}
	if isSameTree(&right, &right) == false {
		t.Errorf("results should be true")
	}

	if isSameTree(&right, &left) == true {
		t.Errorf("results should be false")
	}
}
