package leetcode0086

import "github.com/Ackerr/Algorithms/utils"


func partition(head *utils.ListNode, x int) *utils.ListNode {
	if head == nil {
		return head
	}

	low := &utils.ListNode{Val: -1}
	ans := low
	high := &utils.ListNode{Val: -1}
	target := high
	for head != nil {
		if head.Val < x {
			low.Next = head
			low = low.Next
		} else {
			high.Next = head
			high = high.Next
		}
		head = head.Next
	}
	high.Next = nil
	low.Next = target.Next
	return ans.Next
}
