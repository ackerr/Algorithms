from typing import List


class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        # 自底向上
        length = len(triangle)
        if not length:
            return 0
        for i in range(length - 2, -1, -1):
            for j in range(i+1):
                triangle[i][j] = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
        return triangle[0][0]

    def top_down(self, triangle) -> int:
        # 自顶向下
        length = len(triangle)
        if not length:
            return 0
        for i in range(1, length):
            for j in range(i+1):
                if j == 0:
                    triangle[i][0] += triangle[i-1][0]
                elif i == j:
                    triangle[i][j] += triangle[i-1][j-1]
                else:
                    triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
        return min(triangle[-1])
