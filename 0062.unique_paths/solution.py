import math


class Solution:
    def unique_paths(self, m: int, n: int) -> int:
        """
        到达(m,n) == 到达(m, n-1) + (m-1,n)
        m,或n为1是，路径只会是1
        """
        dp = [1] * n

        for i in range(1, m):
            for j in range(1, n):
                dp[j] += dp[j - 1]

        return dp[-1]

    def arrange(self, m, n):
        """
        从左上角走到右下角需要走 m+n-2 步
        从第一行走到最下行需要走 m - 1 步
        从第一排走到最右排需要走 n - 1 步
        """

        return int(
            math.factorial(m + n - 2) / math.factorial(m - 1) / math.factorial(n - 1)
        )


if __name__ == "__main__":
    assert Solution().unique_paths(3, 3) == 6
    assert Solution().unique_paths(4, 3) == 10

    assert Solution().arrange(3, 2) == 6
    assert Solution().arrange(4, 3) == 10
