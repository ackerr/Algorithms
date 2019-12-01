from typing import List


class Solution:
    def wiggle_max_length(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return len(nums)
        temp = nums[1] - nums[0]
        count = 2 if temp else 1
        for i in range(2, len(nums)):
            diff = nums[i] - nums[i - 1]
            if (diff > 0 and temp <= 0) or (diff < 0 and temp >= 0):
                count += 1
                temp = diff
        return count
