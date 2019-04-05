import collections
from typing import List


class Solution:

    def next_greater_element(self, nums1: List[int], nums2: List[int]) -> List[int]:
        """ use stack """
        values = collections.defaultdict(lambda: -1)
        stack = []
        for num in nums2:
            while stack and stack[-1] < num:
                values[stack.pop()] = num
            stack.append(num)
        return [values[num] for num in nums1]

    def no_use_stack(self, nums1, nums2):
        """ no use stack """
        def find(n, nums):
            if n == nums[-1]:
                return -1
            for m in nums[nums.index(n) + 1:]:
                if m > n:
                    return m
            return -1
        return [find(n, nums2) for n in nums1]
