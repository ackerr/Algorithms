package main

import (
	"fmt"
	"strconv"
)

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	s := node
	var results string
	for s != nil {
		results += strconv.Itoa(s.Val)
		s = s.Next
	}
	return results

}

func deleteDuplicates(head *ListNode) *ListNode {
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

func main() {
	zero := ListNode{Val: 0, Next: nil}
	one := ListNode{Val: 1, Next: &zero}
	two := ListNode{Val: 2, Next: &one}
	fmt.Print(deleteDuplicates(&two))
}
