package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	p := head
	for p != nil && p.Next != nil {
		p.Val, p.Next.Val = p.Next.Val, p.Val
		p = p.Next.Next
	}
	return head
}
