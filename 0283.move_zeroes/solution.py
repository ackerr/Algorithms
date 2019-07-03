class Solution:
    def move_zeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        last = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[i], nums[last] = nums[last], nums[i]
                last += 1
