class Solution:
    def max_profit(self, prices):
        min_price, max_profit = 9999999, 0
        for price in prices:
            if min_price > price:
                min_price = price
            elif max_profit < price - min_price:
                max_profit = price - min_price
        return max_profit
