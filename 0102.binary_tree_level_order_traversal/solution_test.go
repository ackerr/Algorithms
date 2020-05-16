package leetcode0102

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	p := TreeNode{1, nil, nil}
	left := TreeNode{2, &p, &TreeNode{3, nil, nil}}
	right := TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}
	root := TreeNode{4, &left, &right}

	res := [][]int{{4}, {2, 6}, {1, 3, 5, 7}}
	if !reflect.DeepEqual(levelOrder(&root), res) {
		t.Errorf("result should be %+v", res)
	}
}
