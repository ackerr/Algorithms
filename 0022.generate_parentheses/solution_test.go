package main

import (
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	ans := generateParenthesis(2)
	if len(ans) != 2 {
		t.Errorf("resutls error")
	}

	ans = generateParenthesis(3)
	if len(ans) != 5 {
		t.Errorf("resutls error")
	}
}
