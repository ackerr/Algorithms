package leetcode0142

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	fourth := utils.ListNode{Val: 4}
	third := utils.ListNode{Val: 3, Next: &fourth}
	second := utils.ListNode{Val: 2, Next: &third}
	first := utils.ListNode{Val: 1, Next: &second}
	fourth.Next = &second

	if detectCycle(&first) != &second {
		t.Errorf("results error, shoule be %d", second.Val)
	}
}
