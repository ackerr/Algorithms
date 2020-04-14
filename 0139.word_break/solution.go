package main

func wordBreak(s string, wordDict []string) bool {
	b := []byte(s)
	dp := make([]bool, len(b))
	for i := 0; i < len(s); i++ {
		for _, key := range wordDict {
			if i < len(key)-1 {
				continue
			}
			if (i == len(key)-1 || dp[i-len(key)]) && string(b[i-len(key)+1:i+1]) == key {
				dp[i] = true
			}
		}
	}
	return dp[len(s)-1]
}
