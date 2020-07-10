package leetcode0094

import "github.com/Ackerr/Algorithms/utils"

func inorderTraversal(root *utils.TreeNode) []int {
	if root == nil{
		return []int{}
	}
	var stack []*utils.TreeNode
	var ans []int
	for root != nil || len(stack) > 0{
		for root != nil{
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		root = stack[index]
		ans = append(ans, root.Val)
		root = root.Right
		stack = stack[:index]
	}
	return ans
}