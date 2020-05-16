package leetcode0572

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}
	return isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(m *TreeNode, n *TreeNode) bool {
	if m == nil && n == nil {
		return true
	} else if m == nil || n == nil {
		return false
	}
	return m.Val == n.Val && isSameTree(m.Left, n.Left) && isSameTree(m.Right, n.Right)
}
