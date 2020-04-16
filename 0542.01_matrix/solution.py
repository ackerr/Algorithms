import collections
from typing import List


class Solution:
    def update_matrix(self, matrix: List[List[int]]) -> List[List[int]]:
        rows = len(matrix)
        cols = len(matrix[0])
        zero = [(row, col) for row in range(rows) for col in range(cols) if matrix[row][col] == 0]
        ans = [[0] * cols for _ in range(rows)]
        seen = set(zero)
        q = collections.deque(zero)
        while q:
            row, col = q.popleft()
            for r, c in [(row - 1, col), (row + 1, col), (row, col + 1), (row, col - 1)]:
                if 0 <= c < cols and 0 <= r < rows and (r, c) not in seen:
                    ans[r][c] = ans[row][col] + 1
                    q.append((r, c))
                    seen.add((r, c))
        return ans
