package leetcode0199

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	left := utils.TreeNode{
		Val:   2,
		Left:  &utils.TreeNode{Val: 1, Right: &utils.TreeNode{Val: 2}},
		Right: &utils.TreeNode{Val: 3}}
	right := utils.TreeNode{
		Val:   6,
		Left:  &utils.TreeNode{Val: 5},
		Right: &utils.TreeNode{Val: 7, Right: &utils.TreeNode{Val: 10}},
	}
	node := utils.TreeNode{Val: 4, Left: &left, Right: &right}
	if !reflect.DeepEqual(rightSideView(&node), []int{4, 6, 7, 10}) {
		t.Errorf("results should be [4,6,7]")
	}

}
