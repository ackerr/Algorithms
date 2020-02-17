package main

import "fmt"

func reverse(x int) int {
	ans := 0
	for x != 0 {
		pop := x % 10
		// overflow return 0
		if ans > (2<<30)/10 || ans == (2<<30)/10 && pop > 7 {
			return 0
		}

		if ans < (-2<<30)/10 || ans == (-2<<30)/10 && pop < -8 {
			return 0
		}
		ans = ans*10 + pop
		x /= 10
	}
	return ans
}

func main() {
	fmt.Println(reverse(100))
	fmt.Println(reverse(-100))
	fmt.Println(reverse(2<<30 + 1))
	fmt.Println(reverse(-2 << 30))
}
