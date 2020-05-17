package leetcode0082

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	zero := utils.ListNode{Val: 0, Next: nil}
	one := utils.ListNode{Val: 1, Next: &zero}
	two := utils.ListNode{Val: 2, Next: &one}
	node := utils.ListNode{Val: 2, Next: &two}

	if deleteDuplicates(&node).ToString() != "10" {
		t.Errorf("result should be 10")
	}
}
