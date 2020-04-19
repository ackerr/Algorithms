package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"abc", "acb"}
	ans := groupAnagrams(strs)
	results := [][]string{{"abc", "acb"}}
	if !reflect.DeepEqual(ans, results) {
		t.Errorf("results error")
	}
}
