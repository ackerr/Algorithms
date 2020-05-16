package leetcode0096

func numTrees(n int) int {
	m := make([]int, n+1)
	m[0], m[1] = 1, 1

	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			m[i] += m[j-1] * m[i-j]
		}
	}
	return m[n]
}
