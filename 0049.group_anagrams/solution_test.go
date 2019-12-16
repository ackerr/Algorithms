package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"abc", "acb", "aba", "aab"}
	ans := groupAnagrams(strs)
	results := [][]string{{"abc", "acb"}, {"aba", "aab"}}
	if ! reflect.DeepEqual(ans, results) {
		t.Errorf("results error")
	}
}
