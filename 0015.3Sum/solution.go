package main

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] { // distinct
			continue
		}
		for l < r {

			if l > i+1 && nums[l] == nums[l-1] { // distinct
				l++
				continue
			}
			if r < len(nums)-2 && nums[r] == nums[r+1] { // distinct
				r--
				continue
			}
			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return ans
}
