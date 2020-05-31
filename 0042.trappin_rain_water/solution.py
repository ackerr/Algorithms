from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        """ dp: 保存当前点左区间最大值以及右区间最大值 """
        if not height:
            return 0
        length = len(height)
        left_max = []
        right_max = []

        for i, v in enumerate(height):
            value = v if i == 0 else max(v, left_max[i - 1])
            left_max.append(value)

        for i, v in enumerate(height[::-1]):
            value = v if i == 0 else max(v, right_max[i - 1])
            right_max.append(value)

        ans = 0

        for i in range(length):
            ans += min(left_max[i], right_max[length - i - 1]) - height[i]

        return ans
