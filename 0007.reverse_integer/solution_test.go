package main

import "testing"

func TestReverse(t *testing.T) {
	req := []int{
		100,
		-100,
		111,
		-123,
		-2 << 30,
		2<<30 + 10,
	}

	results := []int{
		1,
		-1,
		111,
		-321,
		0,
		0,
	}
	for i, v := range req {
		ans := reverse(v)
		if ans != results[i] {
			t.Errorf("result should be %d, not be %d", results[i], ans)
		}
	}
}
