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

func trap2(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	ans := 0
	leftMax, rightMax := height[0], height[length-1]
	left, right := 0, length-1
	for left < right {
		if height[left] <= height[right] {
			leftMax = utils.Max(leftMax, height[left])
			ans += leftMax - height[left]
			left++
		} else {
			rightMax = utils.Max(rightMax, height[right])
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
