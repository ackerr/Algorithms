class Solution:
    def next_permutation(self, nums):
        l = len(nums) - 2
        while l > 0 and nums[l + 1] <= nums[l]:
            l -= 1

        if l >= 0:
            j = len(nums) - 1
            while j >= 0 and nums[j] <= nums[l]:
                j -= 1
            nums[l], nums[j] = nums[j], nums[l]
        for i in range(0, len(nums) // 2):
            nums[i], nums[-i-1] = nums[-i-1], nums[i]

