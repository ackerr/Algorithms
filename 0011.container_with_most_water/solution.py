from typing import List


class Solution:
    def max_area(self, height: List[int]) -> int:
        """ 双指针，向内移动寻找最大值"""
        max_area = 0
        left, right = 0, len(height) - 1
        while left < right:
            area = min(height[left], height[right]) * (right - left)
            max_area = max(max_area, area)
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return max_area


if __name__ == "__main__":
    assert Solution().max_area([1, 8, 6, 2, 5, 4, 8, 3, 7]) == 49
    assert Solution().max_area([1, 1, 1, 5, 5, 1, 1]) == 6
