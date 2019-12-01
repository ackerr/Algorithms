from typing import List


class Solution:
    def find_max_average(self, nums: List[int], k: int) -> float:
        ans = max_sum = sum(nums[0:k])
        for i in range(k, len(nums)):
            max_sum += nums[i] - nums[i - k]
            ans = max(max_sum, ans)
        return ans / k


if __name__ == "__main__":
    assert Solution().find_max_average([0, 4, 0, 3, 2], 2) == 2.5
