from typing import List


class Solution:
    def count_negatives(self, grid: List[List[int]]) -> int:
        """
        倒序排列，从左下角位置判断，如果该值小于0，说明后面的值都小于0，上移一位
        否则右移一位，直到到达右上角
        """
        if not grid:
            return 0
        m, n = len(grid) - 1, 0
        ans = 0
        while m >= 0 and n < len(grid[0]):
            if grid[m][n] < 0:
                ans += len(grid[0]) - n
                m -= 1
            else:
                n += 1
        return ans
