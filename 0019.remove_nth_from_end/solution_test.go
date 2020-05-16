package leetcode0019

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	fifth := ListNode{Val: 5}
	fourth := ListNode{Val: 4, Next: &fifth}
	third := ListNode{Val: 3, Next: &fourth}
	second := ListNode{Val: 2, Next: &third}
	first := ListNode{Val: 1, Next: &second}
	ans := removeNthFromEnd(&first, 3)

	r := result(ans)
	if r != 12 {
		t.Errorf("result should be 12, not %d", r)
	}

}

func result(l *ListNode) (r int) {
	for l != nil {
		r += l.Val
		l = l.Next
	}
	return
}
