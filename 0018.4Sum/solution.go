package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) (ans [][]int) {
	if len(nums) < 4 {
		return ans
	}
	sort.Ints(nums)
	for m := 0; m < len(nums)-3; m++ {
		if m > 0 && nums[m-1] == nums[m] {
			continue
		}
		for i := m + 1; i < len(nums)-2; i++ {
			if i-m > 1 && nums[i] == nums[i-1] { // distinct
				continue
			}
			l, r := i+1, len(nums)-1
			for l < r {
				if l > i+1 && nums[l] == nums[l-1] { // distinct
					l++
					continue
				}
				if r < len(nums)-2 && nums[r] == nums[r+1] { // distinct
					r--
					continue
				}
				sum := nums[m] + nums[i] + nums[l] + nums[r]

				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					ans = append(ans, []int{nums[m], nums[i], nums[l], nums[r]})
					l++
					r--
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
