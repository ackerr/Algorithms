package main

import (
	"fmt"
	"testing"
)

func TestNumTree(t *testing.T) {
	num := numTrees(3)
	fmt.Println(num)
	if num != 5 {
		t.Errorf("results error")
	}

	num = numTrees(5)
	fmt.Println(num)
	if num != 42 {
		t.Errorf("results error")
	}

}
