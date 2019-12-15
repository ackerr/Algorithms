package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int{
	if a <= b{
		return a
	}
	return b
}

func divide(dividend int, divisor int) int {
	sign := (dividend < 0) == (divisor < 0)
	a, b, res := abs(dividend), abs(divisor), 0
	var i uint
	for i = 31; i >= 0; i--{
		if (a >> i) >= b {
			res += 1 << i
			a -= b << i
		}
	}
	if !sign{
		res = -res
	}
	return min(res, math.MaxInt32)
}
