package leetcode0005

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLen := 1
	start := 0

	dp := make([][]bool, len(s))
	for j := 0; j < len(s); j++ {
		dp[j] = make([]bool, len(s))
		dp[j][j] = true
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			if dp[i][j] {
				curLen := j - i + 1
				if curLen > maxLen {
					maxLen = curLen
					start = i
				}
			}
		}
	}
	return s[start : start+maxLen]
}
