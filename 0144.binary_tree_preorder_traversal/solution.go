package leetcode0144

import "github.com/Ackerr/Algorithms/utils"

func preorderTraversal(root *utils.TreeNode) []int {
	stack := []*utils.TreeNode{root}
	var ans []int
	for len(stack) > 0 {
		index := len(stack) - 1
		node := stack[index]
		stack = stack[:index]
		if node != nil {
			ans = append(ans, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}
	return ans
}
