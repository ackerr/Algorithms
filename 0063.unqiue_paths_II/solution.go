package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	down := len(obstacleGrid)
	right := len(obstacleGrid[0])

	var dp []int
	for i := 0; i < right; i++ {
		dp = append(dp, 0)
	}
	dp[0] = 1
	for i := 0; i < down; i++ {
		for j := 1; j < right; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[right-1]
}

func main() {
	first := []int{0, 0, 1, 0}
	second := []int{0, 0, 0, 0}
	third := []int{0, 1, 0, 0}
	p := [][]int{first, second, third}
	print(uniquePathsWithObstacles(p) == 4)
}
