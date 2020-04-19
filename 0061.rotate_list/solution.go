package main

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
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
