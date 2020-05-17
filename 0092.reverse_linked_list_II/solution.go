package leetcode0092

import "github.com/Ackerr/Algorithms/utils"

func reverseBetween(head *utils.ListNode, m int, n int) *utils.ListNode {
	if head == nil || m > n {
		return head
	}
	ans := &utils.ListNode{}
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
