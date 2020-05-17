package leetcode0328

import "github.com/Ackerr/Algorithms/utils"

func oddEvenList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
