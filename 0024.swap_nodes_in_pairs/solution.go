package leetcode0024

import "github.com/Ackerr/Algorithms/utils"

func swapPairs(head *utils.ListNode) *utils.ListNode {
	p := head
	for p != nil && p.Next != nil {
		p.Val, p.Next.Val = p.Next.Val, p.Val
		p = p.Next.Next
	}
	return head
}
