package leetcode0215


func findKthLargest(nums []int, k int) int {
	length := len(nums)
	left, right := 0, length - 1
	index := length - k
	for {
		p := partition(nums, left, right)
		if index > p{
			left = p + 1
		}else if index < p{
			right = p - 1
		}else{
			return nums[p]
		}
	}
}

func partition(nums []int, left, right int) int {
	// left,mid,right取中值
	mid := left + (right - left) >> 1
	if nums[left] < nums[right]{
		nums[left], nums[right] = nums[right], nums[left]
	}
	if nums[right] < nums[mid]{
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	if nums[left] < nums[mid]{
		nums[mid], nums[left] = nums[left], nums[mid]
	}

	p := nums[left]
	i := left
	for j:=left+1;j<=right;j++{
		if nums[j] < p{
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}
