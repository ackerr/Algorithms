package main

func canThreePartsEqualSum(A []int) bool {
	s := 0
	for _, v := range A {
		s += v
	}
	if s%3 != 0 {
		return false
	}
	target := s / 3
	curSum := 0
	res := 0
	for i := range A {
		curSum += A[i]
		if curSum == target {
			res++
			curSum = 0
		}
	}
	return res > 2
}
