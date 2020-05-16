package leetcode0075

func sortColors(nums []int) {
	var cur, left int
	right := len(nums) - 1
	for cur <= right {
		if nums[cur] == 0 {
			nums[cur], nums[left] = nums[left], nums[cur]
			cur++
			left++
		} else if nums[cur] == 2 {
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		} else {
			cur++
		}
	}
}
