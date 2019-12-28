package main

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	output := nextPermutation([]int{4, 3, 2, 1})
	result := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(output, result) {
		t.Errorf("result should be %d", result)
	}

	output = nextPermutation([]int{5, 3, 4, 2, 3, 4, 2, 1})
	result = []int{5, 3, 4, 2, 4, 1, 2, 3}
	if !reflect.DeepEqual(output, result) {
		t.Errorf("result should be %d", result)
	}

}
