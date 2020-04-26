from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        def dfs(first):
            if first == length:
                ans.append(nums[:])
            for i in range(first, length):
                nums[i], nums[first] = nums[first], nums[i]
                dfs(first + 1)  # dfs
                nums[i], nums[first] = nums[first], nums[i]  # reduction

        ans = []
        length = len(nums)
        dfs(0)
        return ans
