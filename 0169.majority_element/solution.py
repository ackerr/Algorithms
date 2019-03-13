class Solution:

    def majority_element(self, nums):
        total, ans = 0, 0
        for num in nums:
            if total == 0:
                ans = num
            if num == ans:
                total += 1
            else:
                total -= 1
        return ans

    def other_solution(self, nums):
        """ one line"""
        return sorted(nums)[len(nums)//2]
