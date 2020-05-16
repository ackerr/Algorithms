package leetcode0152

import "math"

func maxProduct(nums []int) int {
	ans := math.MinInt32
	currentMax, currentMin := 1, 1
	for _, n := range nums {
		if n < 0 {
			currentMax, currentMin = currentMin, currentMax
		}
		currentMax = max(currentMax*n, n)
		currentMin = min(currentMin*n, n)
		ans = max(currentMax, ans)
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
