package utils

import "strconv"

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

// ToString : join node value to string
func (node *ListNode) ToString() string {
	s := node
	var results string
	for s != nil {
		results += strconv.Itoa(s.Val)
		s = s.Next
	}
	return results

}

// Min : utils.Min(1,2) = 1
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max : utils.Max(1,2) = 2
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abs : utils.Abs(-1) == 1
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
