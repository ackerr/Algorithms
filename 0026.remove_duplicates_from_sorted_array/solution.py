class Solution:
    def remove_duplicates(self, nums):
        total = 0
        if not nums:
            return total
        for index, num in enumerate(nums[:]):
            if index != 0 and num != nums[total]:
                total += 1
                nums[total] = nums[index]
        return total + 1
