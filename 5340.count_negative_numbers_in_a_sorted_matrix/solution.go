package leetcode5340

func countNegatives(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	ans := 0
	m, n := len(grid)-1, 0
	for m > -1 && n < len(grid[0]) {
		if grid[m][n] < 0 {
			ans += len(grid[0]) - n
			m--
		} else {
			n++
		}
	}
	return ans
}
