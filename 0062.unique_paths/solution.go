package main

import "fmt"

func uniquePaths(m int, n int) int {

	var dp []int
	for i := 0; i < n; i++ {
		dp = append(dp, 1)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			fmt.Println(dp)
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

func main() {
	print(uniquePaths(3, 2))
	print(uniquePaths(3, 4))
}
