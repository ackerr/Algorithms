package leetcode0082

import (
	"github.com/Ackerr/Algorithms/utils"
)

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}
