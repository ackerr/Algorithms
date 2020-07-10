package leetcode0094

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestInOrder(t *testing.T) {
	left := utils.TreeNode{
		Val:   3,
		Left:  &utils.TreeNode{Val: 1, Right: &utils.TreeNode{Val: 2}},
		Right: &utils.TreeNode{Val: 4}}
	node := utils.TreeNode{Val: 5, Left: &left, Right: nil}
	answer := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(inorderTraversal(&node), answer) {
		t.Errorf("result should be %+v", answer)
	}
}
