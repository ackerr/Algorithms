from typing import List


class Solution:
    def unique_paths_with_obstacle(self, obstacle_grid: List[List[int]]) -> int:
        """
        二维数组 关键点在于 dp[i][j] = dp[i-1][j] + dp[i][j-1]
        如果遇到障碍1，后续都为0
        """
        if obstacle_grid[0][0]:  # 出发点是障碍
            return 0
        down = len(obstacle_grid)
        right = len(obstacle_grid[0])
        dp = [0] * right
        dp[0] = 1
        for i in range(down):
            for j in range(1, right):
                dp[j] = 0 if obstacle_grid[i][j] else dp[j] + dp[j-1]
        return dp[-1]


if __name__ == '__main__':
    grid = [[0, 0, 0], [0, 1, 0], [0, 0, 0]]
    assert Solution().unique_paths_with_obstacle(grid) == 2

    grid = [[0, 0, 0, 1], [0, 1, 0, 0], [0, 0, 0, 0]]
    assert Solution().unique_paths_with_obstacle(grid) == 3
