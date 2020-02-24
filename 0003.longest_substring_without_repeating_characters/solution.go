package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	mark := make(map[byte]int, len(s))
	ans := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if val, ok := mark[s[right]]; ok {
			left = Max(left, val)
		}
		ans = Max(ans, right-left+1)
		mark[s[right]] = right + 1
	}
	return ans
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}