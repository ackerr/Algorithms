package leetcode0440

import "github.com/Ackerr/Algorithms/utils"

func findKthNumber(n int, k int) int {
	cur := 1
	k--
	for k > 0 {
		step := calStep(n, cur, cur+1)
		if step <= k {
			k -= step
			cur++
		} else {
			k--
			cur *= 10
		}
	}
	return cur
}

func calStep(n, left, right int) int {
	step := 0
	for left <= n {
		step += utils.Min(n+1, right) - left
		left *= 10
		right *= 10
	}
	return step
}
