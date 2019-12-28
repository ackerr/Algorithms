class Solution:
    def next_permutation(self, nums):
        i = len(nums) - 2
        while i >= 0 and nums[i + 1] <= nums[i]:  # 找到第一个不是倒序排列的值
            i -= 1

        if i >= 0:  # 找到第一个比i大的值
            j = len(nums) - 1
            while j >= 0 and nums[j] <= nums[i]:
                j -= 1
            nums[i], nums[j] = nums[j], nums[i]
        left, right = i+1, len(nums) - 1
        while left < right:  # 反转倒序排列的值，形成下一个值，或最小值
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1
