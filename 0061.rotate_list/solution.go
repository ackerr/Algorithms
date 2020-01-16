package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 1
	old_tail := head
	new_tail := head
	for old_tail.Next != nil {
		old_tail = old_tail.Next
		count++
	}
	old_tail.Next = head
	for i := 0; i < count-k%count-1; i++ {
		new_tail = new_tail.Next
	}
	ans := new_tail.Next
	new_tail.Next = nil
	return ans
}
