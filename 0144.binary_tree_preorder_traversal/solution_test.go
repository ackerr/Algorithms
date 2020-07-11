package leetcode0144

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestPreOrder(t *testing.T) {
	left := utils.TreeNode{
		Val:   3,
		Left:  &utils.TreeNode{Val: 1, Right: &utils.TreeNode{Val: 2}},
		Right: &utils.TreeNode{Val: 4}}
	right := utils.TreeNode{
		Val:   8,
		Left:  &utils.TreeNode{Val: 6, Right: &utils.TreeNode{Val: 7}},
		Right: &utils.TreeNode{Val: 9}}
	node := utils.TreeNode{Val: 5, Left: &left, Right: &right}
	answer := []int{5, 3, 1, 2, 4, 8, 6, 7, 9}
	if !reflect.DeepEqual(preorderTraversal(&node), answer) {
		t.Errorf("result should be %+v", answer)
	}
}
