class Solution:
    def remove_element(self, nums: list, val: int) -> int:
        flag = 0
        for index in range(len(nums)):
            if nums[index] != val:
                nums[flag] = nums[index]
                flag += 1
        return flag
