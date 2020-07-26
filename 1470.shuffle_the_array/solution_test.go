package leetcode1470

import (
	"reflect"
	"testing"
)

type item struct {
	input  []int
	n      int
	output []int
}

func TestShuffle(t *testing.T) {
	data := []item{
		{input: []int{2, 5, 1, 3, 4, 7}, n: 3, output: []int{2, 3, 5, 4, 1, 7}},
		{input: []int{1, 2, 3, 4, 5, 6}, n: 3, output: []int{1, 4, 2, 5, 3, 6}},
	}

	for _, item := range data {
		if !reflect.DeepEqual(item.output, shuffle(item.input, item.n)) {
			t.Errorf(
				"shuffle(%+v, %d) should be %+v, not %+v",
				item.input, item.n, item.output, shuffle(item.input, item.n),
			)
		}
	}
}
