package leetcode0004

import (
	"github.com/Ackerr/Algorithms/utils"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m1, m2 := len(nums1), len(nums2)
	if m1 > m2{ // 保证nums1长度较短
		return findMedianSortedArrays(nums2, nums1)
	}
	k := (m1 + m2 + 1) >> 1  // 如果为奇数时，左半边多一个, 偶数时+1不影响
	left, right := 0, m1
	for left < right {
		mid := left + (right-left)>>1
		j := k - mid
		if nums1[mid] < nums2[j - 1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	n1, n2 := left, k - left
	left1, left2 := math.MinInt32, math.MinInt32
	if n1 > 0{
		left1 = nums1[n1-1]
	}
	if n2 > 0{
		left2 = nums2[n2-1]
	}
	left = utils.Max(left1, left2)
	if (m1+m2) % 2 == 1{
		return float64(left)
	}
	right1, right2 := math.MaxInt32, math.MinInt32
	if n1 < m1{
		right1 = nums1[n1]
	}
	if n2 < m2{
		right2 = nums2[n2]
	}
	right = utils.Min(right1, right2)
	return float64(left + right) / 2.0
}
