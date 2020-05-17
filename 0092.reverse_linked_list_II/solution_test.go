package leetcode0092

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	firth := utils.ListNode{5, nil}
	fourth := utils.ListNode{4, &firth}
	third := utils.ListNode{3, &fourth}
	second := utils.ListNode{Val: 2, Next: &third}
	first := utils.ListNode{Val: 1, Next: &second}

	res := []int{1, 5, 4, 3, 2}

	r := reverseBetween(&first, 2, 5)
	for i := 0; i < 5; i++ {
		if r.Val != res[i] {
			t.Errorf("index %d 's value should be %d", i, res[i])
		}
		r = r.Next
	}
}
