package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m > n {
		return head
	}
	ans := &ListNode{}
	ans.Next = head
	p := ans
	for i := 1; i < m; i++ {
		p = p.Next
	}
	head = p.Next
	for i := m; i < n; i++ {
		temp := head.Next
		head.Next = temp.Next
		temp.Next = p.Next
		p.Next = temp
	}
	return ans.Next
}
