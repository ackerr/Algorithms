package leetcode0029

import (
	"github.com/Ackerr/Algorithms/utils"
	"math"
)


func divide(dividend int, divisor int) int {
	sign := (dividend < 0) == (divisor < 0)
	a, b, res := utils.Abs(dividend), utils.Abs(divisor), 0
	// if i -> int build error, i -> uint leetcode error
	//for i = 31; i >= 0; i--{
	//	if (a >> i) >= b {
	//		res += 1 << i
	//		a -= b << i
	//	}
	//}
	for a >= b {
		var x uint = 0
		for a >= (b << (x + 1)) {
			x++
		}
		res += 1 << x
		a -= b << x
	}
	if !sign {
		res = -res
	}
	return utils.Min(res, math.MaxInt32)
}
