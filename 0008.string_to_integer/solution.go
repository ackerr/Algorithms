package leetcode0008

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	sign := 1
	switch str[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
	case '-':
		str = str[1:]
		sign = -1
	case '+':
		str = str[1:]
	default:
		return 0
	}
	for i, v := range str {
		if v < '0' || v > '9' {
			str = str[:i]
			break
		}
	}

	absNum := 0
	for _, b := range str {
		absNum = absNum*10 + int(b-'0')
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && absNum*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * absNum
}
