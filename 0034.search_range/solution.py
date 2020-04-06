from typing import List


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] < target:
                left = mid + 1
            elif nums[mid] > target:
                right = mid - 1
            else:
                left = right = mid
                while left > 0 and nums[left - 1] == target:
                    left -= 1
                while right < len(nums) - 1 and nums[right + 1] == target:
                    right += 1
                return [left, right]
        return [-1, -1]
