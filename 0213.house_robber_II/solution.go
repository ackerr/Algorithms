package main

func rob(nums []int) int {
	length := len(nums)
	if length == 0{
		return 0
	}
	if length < 3{
		return max(nums)
	}
	return max([]int{sub(nums[1:]), sub(nums[:length-1])})
}


func sub(nums []int) int{
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max([]int{nums[0], nums[1]})
	for i := 2;i<len(nums);i++{
		dp[i] = max([]int{dp[i-1], dp[i-2] + nums[i]})
	}
	return max(dp)
}

func max(nums []int) int {
	v := nums[0]
	for _, n := range nums{
		if n > v{
			v = n
		}
	}
	return v
}
