package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	a := ListNode{Val: 1}
	b := ListNode{Val: 2, Next: &a}
	c := ListNode{Val: 5, Next: &b}

	d := ListNode{Val: 2}
	e := ListNode{Val: 5, Next: &a}
	f := ListNode{Val: 0, Next: &b}
	ans := addTwoNumbers(&c, &c)

	actual := []int{}
	for ans != nil {
		actual = append(actual, ans.Val)
		ans = ans.Next
	}
	expext := []int{}
	for f != nil {
		expect = append(expect, f.Val)
		f = f.Next
	}

	if !reflect.deepEqual(actual, f) {
		t.Errorf("result error")
	}
}
