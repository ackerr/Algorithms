package leetcode0023

import (
	"github.com/Ackerr/Algorithms/utils"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	one := &utils.ListNode{Val: 1}
	two := &utils.ListNode{Val: 2}
	three := &utils.ListNode{Val: 3}
	four := &utils.ListNode{Val: 4}
	five := &utils.ListNode{Val: 5}
	six := &utils.ListNode{Val: 6}
	one.Next = six
	two.Next = five
	three.Next = four
	req := []*utils.ListNode{one, two, three}
	if mergeKLists(req).ToString() != "123456" {
		t.Errorf("resutls should be 123456")
	}
}
