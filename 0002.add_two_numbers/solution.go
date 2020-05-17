package leetcode0002

import (
	"github.com/Ackerr/Algorithms/utils"
)

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	ans := &utils.ListNode{Val: -1}
	cur := ans
	flag := 0
	for l1 != nil || l2 != nil || flag == 1 {
		val := flag
		if l1 != nil {
			val += l1.Val
		}
		if l2 != nil {
			val += l2.Val
		}
		flag = val / 10
		cur.Next = &utils.ListNode{Val: val % 10}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return ans.Next
}
