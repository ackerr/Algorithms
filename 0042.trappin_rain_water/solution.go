package leetcode0042

import (
	"github.com/Ackerr/Algorithms/utils"
)

func trap(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	leftMax := make([]int, length)
	leftMax[0] = height[0]
	for i := 1; i < length; i++ {
		leftMax[i] = utils.Max(leftMax[i-1], height[i])

	}
	rightMax := make([]int, length)
	rightMax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = utils.Max(rightMax[i+1], height[i])
	}

	ans := 0
	for i := 0; i < length; i++ {
		ans += utils.Min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}
