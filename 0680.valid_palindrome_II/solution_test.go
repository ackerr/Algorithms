package leetcode0680

import "testing"

func TestValidPalindrome(t *testing.T) {
	req := []string{"abcdca", "abbcbba", "a", "", "aaaaaaaaaabaaa", "abcdefg"}
	res := []bool{true, true, true, true, true, false}

	for i, value := range req {
		if validPalindrome(value) != res[i] {
			t.Errorf("value %s is Palindrome? %t", value, res[i])
		}
	}
}
