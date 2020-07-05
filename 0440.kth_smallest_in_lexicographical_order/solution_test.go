package leetcode0440

import "testing"

type item struct {
	n      int
	k      int
	output int
}

func TestFindKthNumber(t *testing.T) {
	reqs := []item{
		{n: 11, k: 2, output: 10},
		{n: 111, k: 16, output: 12},
		{n: 111, k: 24, output: 2},
		{n: 111, k: 34, output: 29},
	}
	for _, r := range reqs {
		if findKthNumber(r.n, r.k) != r.output {
			t.Errorf("findKthNumber(%d, %d) should be %d", r.n, r.k, r.output)
		}
	}
}
