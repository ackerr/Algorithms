package leetcode0147

import "github.com/Ackerr/Algorithms/utils"

func insertionSortList(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := utils.ListNode{Val: 0}
	ans.Next = head
	for head != nil && head.Next != nil {
		if head.Val <= head.Next.Val {
			head = head.Next
			continue
		}
		pre := &ans
		for pre.Next.Val < head.Next.Val {
			pre = pre.Next
		}
		temp := head.Next
		head.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}
	return ans.Next
}
