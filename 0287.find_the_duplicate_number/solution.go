package leetcode0287

func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 0, 0
	fast = nums[nums[fast]]
	slow = nums[slow]
	for nums[fast] != nums[slow] {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast
}
