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

    def double_pointer(self, height: List[int]) -> int:
        """双指针
        当前值是否能积累雨水, 看当前值左区间最大值和右区间最大值，两者较小的值，与当前值比较
        但是如果左区间最大值，比右区间的某值小，也就不用在意右边的最大值了，例如[1,2,0,3,1,2],
        顺序遍历时，1，2，0都没比最后一位2大, 所以比较是否有积水时，都是看左区间, 反之亦然
        """
        if not height:
            return 0
        left_max, right_max = height[0], height[-1]
        left, right = 0, len(height) - 1
        ans = 0
        while left < right:
            if height[left] <= height[right]:
                left_max = max(left_max, height[left])
                ans += left_max - height[left]
                left += 1
            else:
                right_max = max(right_max, height[right])
                ans += right_max - height[right]
                right -= 1
        return ans
