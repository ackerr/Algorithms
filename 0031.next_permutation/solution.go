package main

func nextPermutation(nums []int) []int {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= i && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	left, right := i+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
