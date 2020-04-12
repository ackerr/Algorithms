from typing import List


class Solution:
    def max_product(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        ans = float("-inf")
        current_max = current_min = 1
        for n in nums:
            if n < 0:
                current_max, current_min = current_min, current_max
            current_max = max(n, current_max * n)
            current_min = min(n, current_min * n)
            ans = max(ans, current_max)
        return ans


if __name__ == '__main__':
    print(Solution().max_product([2, 4, 2, -4]))
    print(Solution().max_product([2, 4, -2, 4]))
    print(Solution().max_product([2, 0, -2, -4]))
    print(Solution().max_product([0, 2]))
    print(Solution().max_product([-3, -4, -5]))
