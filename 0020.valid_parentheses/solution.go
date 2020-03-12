package main

func isValid(s string) bool {
	pair := map[rune]rune{'}': '{', ')': '(', ']': '['}
	c := []rune{}
	for _, i := range s {
		if i == '{' || i == '(' || i == '[' {
			c = append(c, i)
		} else {
			if len(c) == 0 {
				return false
			}
			j := c[len(c)-1]
			if j != pair[i] {
				return false
			}
			c = c[:len(c)-1]
		}
	}
	if len(c) == 0 {
		return true
	}
	return false
}
