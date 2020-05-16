package leetcode0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	down := len(obstacleGrid)
	right := len(obstacleGrid[0])

	obstacleGrid[0][0] = 1

	for i := 1; i < right; i++ {
		if obstacleGrid[0][i] == 1 {
			obstacleGrid[0][i] = 0
		} else {
			obstacleGrid[0][i] = obstacleGrid[0][i-1]
		}
	}

	for i := 1; i < down; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		}
	}
	for i := 1; i < down; i++ {
		for j := 1; j < right; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}

	return obstacleGrid[down-1][right-1]
}

func main() {
	first := []int{0, 0, 1, 0}
	second := []int{0, 0, 0, 0}
	third := []int{0, 1, 0, 0}
	p := [][]int{first, second, third}
	print(uniquePathsWithObstacles(p) == 4)
}
