package leetcode0572

import "testing"

func TestIsSubTree(t *testing.T) {
	p := TreeNode{1, nil, nil}
	left := TreeNode{2, &p, &TreeNode{3, nil, nil}}
	right := TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}
	node := TreeNode{4, &left, &right}
	if isSubtree(&node, &left) == false {
		t.Errorf("results should be true")
	}

	node = TreeNode{2, &p, &TreeNode{3, &p, nil}}
	if isSubtree(&node, &left) == true {
		t.Errorf("results should be false")
	}
}
