from typing import List


class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        queue = []
        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == 2:
                    queue.append((i, j, 0))

        ans = 0
        while queue:
            i, j, ans = queue.pop(0)
            for row, col in ((i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)):
                if 0 <= row < rows and 0 <= col < cols and grid[row][col] == 1:
                    grid[row][col] = 2
                    queue.append((row, col, ans + 1))

        if any(1 in row for row in grid):
            return -1
        return ans


if __name__ == "__main__":
    print(Solution().orangesRotting([[2, 1, 1], [1, 1, 0], [0, 1, 1]]))
