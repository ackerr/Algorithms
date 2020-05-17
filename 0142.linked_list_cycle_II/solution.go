package leetcode0142

import "github.com/Ackerr/Algorithms/utils"

func detectCycle(head *utils.ListNode) *utils.ListNode {
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
