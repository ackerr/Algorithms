from typing import List


class Solution:
    def pivot_index(self, nums: List[int]) -> int:
        total_sum = sum(nums)
        this_sum = 0
        for i in range(len(nums)):
            if total_sum - nums[i] == this_sum:
                return i
            this_sum += nums[i]
            total_sum -= nums[i]
        return -1


if __name__ == '__main__':
    assert Solution().pivot_index([-1, -1, 0, 1, 1, 0]) == 5
