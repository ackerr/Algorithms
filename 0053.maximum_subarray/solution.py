class Solution:

    def max_sub_array(self, nums):
        total, res = 0, nums[0]
        for num in nums:
            if total > 0:
                total += num
            else:
                total = num
            res = max(total, res)
        return res
