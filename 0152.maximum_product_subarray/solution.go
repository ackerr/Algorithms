package leetcode0152

import (
	"github.com/Ackerr/Algorithms/utils"
	"math"
)

func maxProduct(nums []int) int {
	ans := math.MinInt32
	currentMax, currentMin := 1, 1
	for _, n := range nums {
		if n < 0 {
			currentMax, currentMin = currentMin, currentMax
		}
		currentMax = utils.Max(currentMax*n, n)
		currentMin = utils.Min(currentMin*n, n)
		ans = utils.Max(currentMax, ans)
	}
	return ans
}
