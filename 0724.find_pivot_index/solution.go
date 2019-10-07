package main


func pivotIndex(nums []int) int {
	var totalSum int
	for _, v := range nums {totalSum += v}
	partSum := 0
	for i := 0; i < len(nums); i++ {
		if totalSum-nums[i] == partSum {
			return i
		}
		totalSum -= nums[i]
		partSum += nums[i]
	}
	return -1
}
