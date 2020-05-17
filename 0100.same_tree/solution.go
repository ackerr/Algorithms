package leetcode0100

import "github.com/Ackerr/Algorithms/utils"

func isSameTree(m *utils.TreeNode, n *utils.TreeNode) bool {
	if m == nil && n == nil {
		return true
	} else if m == nil || n == nil {
		return false
	}
	return m.Val == n.Val && isSameTree(m.Left, n.Left) && isSameTree(m.Right, n.Right)
}
