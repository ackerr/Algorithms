package main

import "testing"

func TestThreePartsEqualSum(t *testing.T) {
	req := []int{24, -4, -5, -12, 5, 16, -12, 22, 2}
	if canThreePartsEqualSum(req) {
		t.Errorf("result should be false")
	}
	req = []int{12, -4, 16, -5, 9, -3, 3, 8, 0}
	if !canThreePartsEqualSum(req) {
		t.Errorf("result should be true")
	}
}
