package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	req := [][]string{
		[]string{"flower", "flow", "flight"},
		[]string{"a", "as", "ass"},
		[]string{"bab", "bcc"},
		[]string{"aca", "cba"},
	}

	res := []string{"fl", "a", "b", ""}
	for i, v := range req {
		r := longestCommonPrefix(v)
		if r != res[i] {
			t.Errorf("result shoude be %s, not %s", res[i], r)
		}
	}

}
