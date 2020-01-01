from typing import List


class Solution:
    def min_path_sum(self, grid: List[List[int]]) -> int:
        """
        到达一个节点只有两种情况，上面的节点和左边的节点，只需要算个较小值，就是最小路径
        f(i, j) = min(f(i-1, j), f(i, j-1)) + grid(i, j)

        还可以直接修改原数组，空间复杂度为O(1)
        """
        ans = [0] * len(grid[0])
        ans[0] = grid[0][0]
        for i in range(1, len(ans)):
            ans[i] += ans[i - 1] + grid[0][i]
        for i in range(1, len(grid)):
            for j in range(len(grid[0])):
                if j == 0:
                    ans[j] += grid[i][j]
                else:
                    ans[j] = min(ans[j], ans[j - 1]) + grid[i][j]
        return ans[-1]


if __name__ == '__main__':
    assert Solution().min_path_sum([[1, 3, 1], [1, 5, 1], [4, 2, 1]]) == 7
