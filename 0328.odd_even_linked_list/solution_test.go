package leetcode0328

import (
	"testing"
)

func TestOddEvenList(t *testing.T) {
	firth := ListNode{Val: 5}
	fourth := ListNode{Val: 4, Next: &firth}
	third := ListNode{Val: 3, Next: &fourth}
	second := ListNode{Val: 2, Next: &third}
	first := ListNode{Val: 1, Next: &second}
	result := oddEvenList(&first)
	var r []int
	for result != nil {
		r = append(r, result.Val)
		result = result.Next
	}
}
