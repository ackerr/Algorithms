package main

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
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
