from typing import List


class Solution:
    def unique_paths_with_obstacle(self, obstacle_grid: List[List[int]]) -> int:
        """
        二维数组 关键点在于 obstacle_grid[i][j] = obstacle_grid[i-1][j] + obstacle_grid[i][j-1]
        如果遇到障碍1，后续都为0
        """
        if obstacle_grid[0][0]:  # 出发点是障碍
            return 0

        down = len(obstacle_grid)
        right = len(obstacle_grid[0])

        obstacle_grid[0][0] = 1

        # 判断第一行的情况
        for i in range(1, right):
            obstacle_grid[0][i] = obstacle_grid[0][i - 1] if obstacle_grid[0][i] == 0 else 0

        # 判断第一列的情况
        for i in range(1, down):
            obstacle_grid[i][0] = obstacle_grid[i - 1][0] if obstacle_grid[i][0] == 0 else 0

        # 从(1,1)开始将上和左的次数相加
        for i in range(1, down):
            for j in range(1, right):
                if obstacle_grid[i][j] == 1:
                    obstacle_grid[i][j] = 0
                else:
                    obstacle_grid[i][j] = obstacle_grid[i][j - 1] + obstacle_grid[i - 1][j]

        return obstacle_grid[down - 1][right - 1]


if __name__ == "__main__":
    grid = [[0, 0, 0], [0, 1, 0], [0, 0, 0]]
    assert Solution().unique_paths_with_obstacle(grid) == 2

    grid = [[0, 0, 0, 1], [0, 1, 0, 0], [0, 0, 0, 0]]
    assert Solution().unique_paths_with_obstacle(grid) == 3
