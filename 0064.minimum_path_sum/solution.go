package main

func minPathSum(grid [][]int) int {
	length := len(grid[0])
	ans := make([]int, length)
	ans[0] = grid[0][0]
	for i := 1; i < length; i++ {
		ans[i] = ans[i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		for j := 0; j < length; j++ {
			if j == 0 {
				ans[j] += grid[i][j]
			} else {
				ans[j] = min(ans[j], ans[j-1]) + grid[i][j]
			}
		}
	}
	return ans[length-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
