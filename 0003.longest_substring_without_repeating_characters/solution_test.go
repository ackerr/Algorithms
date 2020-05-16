package leetcode0003

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	value := map[string]int{
		"abcdcba":  4,
		"ababccba": 3,
		" ":        1,
		"":         0,
	}
	for k, v := range value {
		result := lengthOfLongestSubstring(k)
		if result != v {
			t.Errorf("result error")
		}
	}
}
