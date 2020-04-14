from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        length = len(nums)
        if length == 0:
            return 0
        if length < 3:
            return max(nums)

        def _rob(m):
            dp = [0] * len(m)
            dp[0] = m[0]
            dp[1] = max(m[0], m[1])
            for i in range(2, len(m)):
                dp[i] = max(dp[i-1], dp[i-2] + m[i])
            return max(dp)

        # 因为是环形排列，所以取移除第一个和最后一个的最大值即可
        return max(_rob(nums[1:]), _rob(nums[:-1]))
