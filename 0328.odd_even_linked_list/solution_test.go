package leetcode0328

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	firth := utils.ListNode{Val: 5}
	fourth := utils.ListNode{Val: 4, Next: &firth}
	third := utils.ListNode{Val: 3, Next: &fourth}
	second := utils.ListNode{Val: 2, Next: &third}
	first := utils.ListNode{Val: 1, Next: &second}
	result := oddEvenList(&first)
	var r []int
	for result != nil {
		r = append(r, result.Val)
		result = result.Next
	}
}
