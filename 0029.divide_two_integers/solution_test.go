package main

import "testing"

func TestDivide(t *testing.T) {
	dividend := 100
	divisor := 7
	res := divide(dividend, divisor)
	if res != 14 {
		t.Errorf("results error")
	}
}
