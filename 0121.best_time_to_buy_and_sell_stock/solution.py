class Solution:
    def max_profit(self, prices):
        min_price, max_profit = 1 << 31, 0
        for price in prices:
            min_price = min(price, price)
            max_profit = max(max_profit, price - min_price)
        return max_profit
