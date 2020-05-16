package leetcode0120

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	}
	for i := length - 2; i > -1; i-- {
		for j := 0; j < i+1; j++ {
			if triangle[i+1][j] < triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}
		}
	}
	return triangle[0][0]
}
