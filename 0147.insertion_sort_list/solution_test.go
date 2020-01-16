package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	a := ListNode{Val: 3}
	b := ListNode{Val: 4, Next: &a}
	c := ListNode{Val: 5, Next: &b}
	d := ListNode{Val: 6, Next: &c}
	sortLinkList := insertionSortList(&d)
	result := make([]int, 0)
	for sortLinkList != nil {
		result = append(result, sortLinkList.Val)
		sortLinkList = sortLinkList.Next
	}
	fmt.Println(result)
	correct := []int{3, 4, 5, 6}
	if !reflect.DeepEqual(result, correct) {
		t.Errorf("result error, correct result should be %d", correct)

	}

}
