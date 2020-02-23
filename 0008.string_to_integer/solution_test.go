package main

import "testing"

func TestAtoi(t *testing.T) {
	req := []string{
		"    -41",
		"+    -",
		"  hello world  -41 hi",
		"words and 987",
		"2147483648",
	}

	results := []int{
		-41,
		0,
		0,
		0,
		2147483647,
	}

	for i, v := range req {
		r := myAtoi(v)
		if r != results[i] {
			t.Errorf("result should be %d, not %d", results[i], r)
		}
	}
}
