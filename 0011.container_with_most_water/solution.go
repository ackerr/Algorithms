package main

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	maxArea := 0
	var left, right = 0, len(height) - 1
	for left < right {
		area := Min(height[left], height[right]) * (right - left)
		maxArea = Max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 1, 1, 5, 5, 1}))
}
