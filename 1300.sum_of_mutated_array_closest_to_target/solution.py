from typing import List


class Solution:
    def findBestValue(self, arr: List[int], target: int) -> int:
        left = 0
        right = max(arr)
        while left <= right:
            mid = left + ((right - left) >> 1)
            cur = sum(mid if i > mid else i for i in arr)
            if cur < target:
                left = mid + 1
            elif cur > target:
                right = mid - 1
            else:
                return mid
        ans1 = sum(left if i > left else i for i in arr)
        ans2 = sum(right if i > right else i for i in arr)
        return right if abs(ans1 - target) >= abs(ans2 - target) else left
