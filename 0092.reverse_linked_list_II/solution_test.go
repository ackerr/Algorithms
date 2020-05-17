package leetcode0092

import "testing"

func TestReverseBetween(t *testing.T) {
	firth := ListNode{5, nil}
	fourth := ListNode{4, &firth}
	third := ListNode{3, &fourth}
	second := ListNode{Val: 2, Next: &third}
	first := ListNode{Val: 1, Next: &second}

	res := []int{1,5,4,3,2}

	r := reverseBetween(&first, 2, 5)
	for i:=0;i<5;i++{
		if r.Val != res[i]{
			t.Errorf("index %d 's value should be %d", i, res[i])
		}
		r = r.Next
	}
}