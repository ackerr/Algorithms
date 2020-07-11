package leetcode0145

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestPostOrder(t *testing.T) {
	left := utils.TreeNode{
		Val:   3,
		Left:  &utils.TreeNode{Val: 1, Right: &utils.TreeNode{Val: 2}},
		Right: &utils.TreeNode{Val: 4}}
	right := utils.TreeNode{
		Val:   8,
		Left:  &utils.TreeNode{Val: 6, Right: &utils.TreeNode{Val: 7}},
		Right: &utils.TreeNode{Val: 9}}
	node := utils.TreeNode{Val: 5, Left: &left, Right: &right}
	answer := []int{2, 1, 4, 3, 7, 6, 9, 8, 5}
	if !reflect.DeepEqual(postorderTraversal(&node), answer) {
		t.Errorf("result should be %+v", answer)
	}
}
