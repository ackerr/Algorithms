package leetcode0098

import "github.com/Ackerr/Algorithms/utils"

func isValidBST(root *utils.TreeNode) bool {
	return dfs(root, -1<<63, 1<<63-1)
}

func dfs(node *utils.TreeNode, low, high int) bool {
	if node == nil {
		return true
	}
	if node.Val <= low || node.Val >= high {
		return false
	}
	if !dfs(node.Left, low, node.Val) {
		return false
	}
	if !dfs(node.Right, node.Val, high) {
		return false
	}
	return true
}
