package main

import "testing"

func TestIsValid(t *testing.T) {
	req := []string{"()[]{}", "((()))", "[][)]", ")(())"}
	res := []bool{true, true, false, false}
	for i, r := range req {
		if isValid(r) != res[i] {
			t.Errorf("result should be %t", res[i])
		}
	}

}
