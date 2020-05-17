package leetcode0019

import "github.com/Ackerr/Algorithms/utils"

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	if head == nil {
		return nil
	}
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
