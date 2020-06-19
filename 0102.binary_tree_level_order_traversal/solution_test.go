package leetcode0102

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {

	left := utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 3}}
	right := utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 7}}
	node := utils.TreeNode{Val: 4, Left: &left, Right: &right}

	res := [][]int{{4}, {2, 6}, {1, 3, 5, 7}}
	if !reflect.DeepEqual(levelOrder(&node), res) {
		t.Errorf("result should be %+v", res)
	}
}
