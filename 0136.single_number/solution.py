class Solution:
    def single_number(self, nums):
        """
        0^1 = 1
        1^1 = 0
        """
        value = nums[0]
        for num in nums[1:]:
            value ^= num
        return value

    def other_solution(self, nums):
        """
        [1,1,2,2,3] + [3] == [1,2,3] * 2
        """
        return sum(set(nums)) * 2 - sum(nums)
