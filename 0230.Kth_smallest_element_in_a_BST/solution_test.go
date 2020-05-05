package main

import "testing"

func TestKthSmallest(t *testing.T) {
	left := TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	right := TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}
	node := TreeNode{4, &left, &right}
	if kthSmallest(&node, 4) != 4 {
		t.Errorf("4th smallest value should be 4")
	}
}
