package leetcode1300

import "testing"

type item struct {
	input  []int
	target int
	output int
}

func TestFindBestValue(t *testing.T) {
	items := []item{
		{
			input:  []int{60864, 25176, 27249, 21296, 20204},
			target: 56803,
			output: 11361,
		},
		{
			input:  []int{2, 3, 5},
			target: 10,
			output: 5,
		},
		{
			input:  []int{4, 9, 3},
			target: 10,
			output: 3,
		},
	}

	for _, i := range items {
		if findBestValue(i.input, i.target) != i.output {
			t.Errorf("findBestValue(%+v, %d) should be %d", i.input, i.target, i.output)
		}
	}
}
