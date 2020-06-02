package leetcode0146

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1{
		t.Errorf("results error")
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Errorf("results error")
	}
	cache.Put(4, 4)

	if cache.Get(1) != -1 {
		t.Errorf("results error")
	}

	if cache.Get(3) != 3 {
		t.Errorf("results error")
	}

	if cache.Get(4) != 4 {
		t.Errorf("results error")
	}
}