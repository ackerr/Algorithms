package main

import "testing"

func TestIsSameTree(t *testing.T) {
	p := TreeNode{1, nil, nil}
	left := TreeNode{2, &p, &TreeNode{3, nil, nil}}
	right := TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}
	if isSameTree(&right, &right) == false {
		t.Errorf("results should be true")
	}

	if isSameTree(&right, &left) == true {
		t.Errorf("results should be false")
	}
}
