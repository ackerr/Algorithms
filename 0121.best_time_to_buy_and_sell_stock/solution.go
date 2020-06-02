package leetcode0121

import "github.com/Ackerr/Algorithms/utils"

func maxProfit(prices []int) int {
	minValue, maxProfit := 1<<31, 0
	for _, price := range prices {
		minValue = utils.Min(price, minValue)
		maxProfit = utils.Max(maxProfit, price-minValue)
	}
	return maxProfit
}
