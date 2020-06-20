from random import random
from typing import List


class Solution:
    def pivot(self, nums, left, right):
        v = random.randint(left, right)
        p = nums[v]
        j = left
        for i in range(left + 1, right + 1):
            if nums[i] < p:
                j += 1
                nums[i], nums[j] = nums[j], nums[i]
        nums[v], nums[j] = nums[j], nums[v]
        return j

    def find_Kth_largest(self, nums: List[int], k: int) -> int:
        """
        第k大值
        1.通过快排倒序排好数值,取第k-1个值, 时间复杂度为排序的复杂度O(nlogn)
        2.调整快排, 根据pivot, 分为左右两边，只对影响到len(nums)-k值的区间进行排序, 时间复杂度为排序的复杂度O(n)
        """
        length = len(nums)
        index = length - k
        left, right = 0, length - 1
        while True:
            p = self.pivot(nums, left, right)
            if p > index:
                right -= 1
            elif p < index:
                left += 1
            else:
                return nums[p]
