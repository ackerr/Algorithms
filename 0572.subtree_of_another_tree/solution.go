package leetcode0572

import "github.com/Ackerr/Algorithms/utils"

func isSubtree(s *utils.TreeNode, t *utils.TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}
	return isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(m *utils.TreeNode, n *utils.TreeNode) bool {
	if m == nil && n == nil {
		return true
	} else if m == nil || n == nil {
		return false
	}
	return m.Val == n.Val && isSameTree(m.Left, n.Left) && isSameTree(m.Right, n.Right)
}
