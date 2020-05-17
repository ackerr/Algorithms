package leetcode0061

import "github.com/Ackerr/Algorithms/utils"

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 1
	oldTail := head
	newTail := head
	for oldTail.Next != nil {
		oldTail = oldTail.Next
		count++
	}
	oldTail.Next = head
	for i := 0; i < count-k%count-1; i++ {
		newTail = newTail.Next
	}
	ans := newTail.Next
	newTail.Next = nil
	return ans
}
