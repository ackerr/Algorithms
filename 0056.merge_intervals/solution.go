package leetcode0056

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	// sort
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	for _, interval := range intervals {
		last := len(ans) - 1
		if last == -1 || ans[last][1] < interval[0] {
			ans = append(ans, interval)
		} else {
			ans[last][1] = max(ans[last][1], interval[1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
