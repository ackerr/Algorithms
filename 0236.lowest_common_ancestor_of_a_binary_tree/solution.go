package leetcode0236

import "github.com/Ackerr/Algorithms/utils"


func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
