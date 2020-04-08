package main

import "testing"

func TestWordBreak(t *testing.T) {
	wordDict := []string{"word", "dog", "tree"}
	values := map[string]bool{
		"wordd": false,
		"worddog": true,
		"treeworddog": true,
	}
	for k, v := range values{
		if wordBreak(k, wordDict) != v{
			t.Errorf("wordBreak %s should be %t", k, v)
		}
	}
}