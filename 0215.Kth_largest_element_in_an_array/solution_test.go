package leetcode0215

import "testing"

type item struct {
	input []int
	target int
	output int
}

func TestKthLargestElement(t *testing.T) {
	req := []item{
		{input: []int{1,3,2,4,4,5}, target: 1, output: 5},
		{input: []int{1,3,2,4,4,5}, target: 2, output: 4},
		{input: []int{1,3,2,4,4,5}, target: 3, output: 4},
		{input: []int{1,3,2,4,4,5}, target: 4, output: 3},
	}
	for _, r := range req{
		if findKthLargest(r.input, r.target) != r.output{
			t.Errorf("findKthLargest(%+v, %d) should be %d", r.input, r.target, r.output)
		}
	}
}