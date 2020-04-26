package main

func permute(nums []int) [][]int {
	return dfs(0, len(nums), nums)
}

func dfs(first, length int, nums []int) (ans [][]int) {
	if first == length{
		ans = append(ans, nums)
	}
	for i := first;i<length;i++{
		nums[i], nums[first] =nums[i], nums[first]
		dfs(first+1, length, nums)
		nums[i], nums[first] =nums[i], nums[first]
	}
	return ans
}
