package leetcode0142

import "testing"

func TestDetectCycle(t *testing.T) {
	fourth := ListNode{Val: 4}
	third := ListNode{Val: 3, Next: &fourth}
	second := ListNode{Val: 2, Next: &third}
	first := ListNode{Val: 1, Next: &second}
	fourth.Next = &second

	if detectCycle(&first) != &second {
		t.Errorf("results error, shoule be %d", second.Val)
	}
}
