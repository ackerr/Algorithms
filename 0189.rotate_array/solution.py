from typing import List


class Solution:

    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        k %= n
        nums[:] = nums[n - k:] + nums[:n - k]

    def no_extra_space(self, nums, k):
        """
        k =3
        nums = [1, 2, 3, 4, 5]

        1. -> [2, 1, 3, 4, 5]
        2. -> [2, 1, 5, 4, 3]
        3. -> [3, 4, 5, 1, 2]

        """
        def reverse(values: List, left: int, right: int):
            while left < right:
                nums[left], nums[right] = nums[right], nums[left]
                left += 1
                right -= 1

        n = len(nums)
        k %= n
        reverse(nums, 0, n - 1)
        reverse(nums, 0, k - 1)
        reverse(nums, k, n - 1)
        return nums
