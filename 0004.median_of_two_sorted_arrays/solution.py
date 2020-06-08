from typing import List


class Solution:
    def find_median_sorted_arrays(self, nums1: List[int], nums2: List[int]) -> float:
        """ O(log(n+m)) """
        n1, n2 = len(nums1), len(nums2)
        if n1 > n2:  # 保证nums1长度较小
            return self.find_median_sorted_arrays(nums2, nums1)
        k = (n1 + n2 + 1) // 2
        left, right = 0, n1
        while left < right:
            mid = left + (right - left) // 2
            m2 = k - mid
            if nums1[mid] < nums2[m2 - 1]:
                left = mid + 1
            else:
                right = mid
        # 中位数在两个数组的位置
        m1, m2 = left, k - left
        c1 = max(nums1[m1 - 1] if m1 > 0 else float("-inf"), nums2[m2 - 1] if m2 > 0 else float("-inf"))
        if (n1 + n2) % 2 == 1:
            return c1
        c2 = min(nums1[m1] if m1 < n1 else float("inf"), nums2[m2] if m2 < n2 else float("inf"))
        return (c1 + c2) / 2

    def other_solution(self, nums1, nums2):
        """ 双指针， O(m+n), 两个数组的第k大值 """
        m = len(nums1)
        n = len(nums2)
        length = m + n
        left = right = 0
        index1 = index2 = 0
        for i in range(length // 2 + 1):
            left = right
            if index1 < m and (index2 >= n or nums1[index1] < nums2[index2]):
                right = nums1[index1]
                index1 += 1
            else:
                right = nums2[index2]
                index2 += 1
        return (left + right) / 2 if length % 2 == 0 else right
