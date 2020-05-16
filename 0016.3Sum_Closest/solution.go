package leetcode0016

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	if length < 3 {
		return 0
	}
	ans := nums[0] + nums[1] + nums[2]
	if length == 3 {
		return ans
	}
	for i := 0; i < length-2; i++ {
		l, r := i+1, length-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if abs(s-target) < abs(ans-target) {
				ans = s
			}
			if s < target {
				l++
			} else if s > target {
				r--
			} else {
				return target
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
