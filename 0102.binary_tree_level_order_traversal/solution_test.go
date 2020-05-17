package leetcode0102

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	p := utils.TreeNode{1, nil, nil}
	left := utils.TreeNode{2, &p, &utils.TreeNode{3, nil, nil}}
	right := utils.TreeNode{6, &utils.TreeNode{5, nil, nil}, &utils.TreeNode{7, nil, nil}}
	root := utils.TreeNode{4, &left, &right}

	res := [][]int{{4}, {2, 6}, {1, 3, 5, 7}}
	if !reflect.DeepEqual(levelOrder(&root), res) {
		t.Errorf("result should be %+v", res)
	}
}
