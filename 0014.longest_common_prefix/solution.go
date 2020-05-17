package leetcode0014

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	l := min(strs)
	for i := l; i > 0; i-- {
		if isPrefix(strs, l) {
			break
		}
		l--
	}
	print(l)
	return strs[0][:l]
}

func isPrefix(strs []string, l int) bool {
	prefix := strs[0][:l]
	for _, s := range strs {
		fmt.Println(s[:l], prefix)
		if s[:l] != prefix {
			return false
		}
	}
	return true
}

func min(strs []string) (v int) {
	v = len(strs[0])
	for _, s := range strs {
		if len(s) < v {
			v = len(s)
		}
	}
	return v
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"aca", "cba"}))
}
