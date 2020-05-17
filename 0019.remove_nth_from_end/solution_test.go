package leetcode0019

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	fifth := utils.ListNode{Val: 5}
	fourth := utils.ListNode{Val: 4, Next: &fifth}
	third := utils.ListNode{Val: 3, Next: &fourth}
	second := utils.ListNode{Val: 2, Next: &third}
	first := utils.ListNode{Val: 1, Next: &second}
	ans := removeNthFromEnd(&first, 3)

	r := result(ans)
	if r != 12 {
		t.Errorf("result should be 12, not %d", r)
	}

}

func result(l *utils.ListNode) (r int) {
	for l != nil {
		r += l.Val
		l = l.Next
	}
	return
}
