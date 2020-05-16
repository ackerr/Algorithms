package leetcode0098

import "testing"

func TestIsValidBST(t *testing.T) {

	p := TreeNode{1, nil, nil}
	left := TreeNode{2, &p, &TreeNode{3, nil, nil}}
	right := TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}
	node := TreeNode{4, &left, &right}

	if !isValidBST(&node) {
		t.Errorf("%+v is a BST", node)
	}

	p.Val = 2
	if isValidBST(&node) {
		t.Errorf("%+v is not a BST", node)
	}

}
