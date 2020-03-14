package main

import "fmt"

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	dp := make([]int, length)

	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max([]int{dp[i], dp[j] + 1})
			}
		}
		fmt.Println(dp)
	}
	return max(dp)
}

func max(nums []int) int {
	value := -1
	for _, n := range nums {
		if value < n {
			value = n
		}
	}
	return value
}
