package utils

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Min : find the min value
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max : find the max value
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
