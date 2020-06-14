package leetcode1300

import "github.com/Ackerr/Algorithms/utils"

func findBestValue(arr []int, target int) int {
	left := 0
	right := utils.MaxList(arr)
	for left <= right {
		mid := left + ((right - left) >> 1)
		cur := targetSum(arr, mid)
		if cur > target {
			right = mid - 1
		} else if cur < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	ansLeft := targetSum(arr, left)
	ansRight := targetSum(arr, right)
	if utils.Abs(ansLeft-target) >= utils.Abs(ansRight-target) {
		return right
	}
	return left
}

func targetSum(array []int, target int) int {
	sum := 0
	for _, v := range array {
		if v > target {
			sum += target
		} else {
			sum += v
		}
	}
	return sum
}
