package leetcode0145

import "github.com/Ackerr/Algorithms/utils"

func postorderTraversal(root *utils.TreeNode) []int {
	return dfs(root, []int{})
}

func dfs(node *utils.TreeNode, ans []int) []int {
	if node == nil {
		return ans
	}
	ans = dfs(node.Left, ans)
	ans = dfs(node.Right, ans)
	ans = append(ans, node.Val)
	return ans
}
