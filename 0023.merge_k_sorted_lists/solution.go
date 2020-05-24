package leetcode0023

import "github.com/Ackerr/Algorithms/utils"

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
	k := len(lists)
	if k == 0 || lists == nil {
		return nil
	}
	return merge(lists, 0, k-1)
}

func merge(lists []*utils.ListNode, start, end int) *utils.ListNode {
	if start == end {
		return lists[start]
	}
	if start > end {
		return nil
	}
	// 归并排序
	mid := start + (end-start)/2
	left := merge(lists, start, mid)
	right := merge(lists, mid+1, end)
	return mergeTwoSortedLists(left, right)
}

func mergeTwoSortedLists(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	ans := &utils.ListNode{Val: 0}
	cur := ans
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return ans.Next
}
