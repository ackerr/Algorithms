package leetcode0005

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := []string{
		"aaabaaabaaaa",
		"aaaaaaaaaaaa",
		"abcdeafedcba",
	}
	r := []string{
		"aaabaaabaaa",
		"aaaaaaaaaaaa",
		"a",
	}
	for i, v := range s {
		result := longestPalindrome(v)
		if result != r[i] {
			t.Errorf("result should be %s, not %s", r[i], v)
		}
	}
}
