package leetcode0070

import "testing"

func TestClimbingStairs(t *testing.T) {
	res := map[int]int{
		5:  8,
		10: 89,
		20: 10946,
		30: 1346269,
		40: 165580141,
		50: 20365011074,
	}

	for k, v := range res {
		if climbStairs(k) != v {
			t.Errorf("climbStairs(%d) should be %d", k, v)
		}
	}
}
