package leetcode0830

func largeGroupPositions(S string) [][]int {
	var ans [][]int
	var i int
	for j := 0; j < len(S); j++ {
		if j == len(S)-1 || S[j] != S[j+1] {
			if j-i > 1 {
				ans = append(ans, []int{i, j})
			}
			i = j + 1
		}
	}
	return ans
}
