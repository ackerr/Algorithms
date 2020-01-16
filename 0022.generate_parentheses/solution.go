package main

import "fmt"

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	backtrack("", n, n, &ans)
	return ans
}

func backtrack(s string, left, right int, ans *[]string) {
	fmt.Println(left, right)
	if left == 0 { // s中没左括号了，其余的都是不匹配的值
		for i := 0; i < right; i++ {
			s += ")"
		}
		*ans = append(*ans, s)
		return
	}
	backtrack(s+"(", left-1, right, ans)
	if left < right {
		backtrack(s+")", left, right-1, ans)
	}
}

func main() {
	generateParenthesis(3)
}
