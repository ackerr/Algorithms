package leetcode0091

import "testing"

func TestNumDecoding(t *testing.T) {
	n := numDecodings("12")
	if n != 2 {
		t.Errorf("result should be 2")
	}

	n = numDecodings("17")
	if n != 2 {
		t.Errorf("result should be 2")
	}
}
