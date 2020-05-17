package leetcode0230

import "github.com/Ackerr/Algorithms/utils"

func kthSmallest(root *utils.TreeNode, k int) int {
	stack := make([]*utils.TreeNode, 0)
	kthVal := 0
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			k--
			root = stack[len(stack)-1] // python list.pop(）
			kthVal = root.Val
			stack = stack[:len(stack)-1]
			root = root.Right
		}
		if k == 0 {
			break
		}
	}
	return kthVal
}
