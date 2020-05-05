package main

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	p := TreeNode{1, nil, nil}
	left := TreeNode{2, &p, &TreeNode{3, nil, nil}}
	right := TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}
	node := TreeNode{4, &left, &right}
	if lowestCommonAncestor(&node, &p, &right) != &node {
		t.Errorf("result error")
	}
}
