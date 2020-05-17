package leetcode0011

import "github.com/Ackerr/Algorithms/utils"


func maxArea(height []int) int {
	maxArea := 0
	var left, right = 0, len(height) - 1
	for left < right {
		area := utils.Min(height[left], height[right]) * (right - left)
		maxArea = utils.Max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
