package main

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	ans := letterCombinations("23")
	if !reflect.DeepEqual(ans, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}){
		t.Errorf("results should be  %s", ans)
	}
}
