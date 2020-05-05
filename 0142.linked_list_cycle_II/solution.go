package main

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	// Check is cycle linked list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if slow == nil || slow.Next == nil {
		return nil
	}
	for head != slow {
		head, slow = head.Next, slow.Next
	}
	return head
}
