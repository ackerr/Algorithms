package main

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(m *TreeNode, n *TreeNode) bool {
	if m == nil && n == nil {
		return true
	} else if m == nil || n == nil {
		return false
	}
	return m.Val == n.Val && isSameTree(m.Left, n.Left) && isSameTree(m.Right, n.Right)
}
