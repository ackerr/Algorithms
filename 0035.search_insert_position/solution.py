class Solution:
    def search_insert(self, nums, target):
        for index, num in enumerate(nums):
            if target <= num:
                return index
        return len(nums)
