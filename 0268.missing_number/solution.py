class Solution:
    def missing_umber(self, nums: List[int]) -> int:
        n = len(nums)
        return n * (n - 1) / 2 - sum(nums)
