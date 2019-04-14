from typing import List

class Solution:
    def rob(self, nums: List[int]) -> int:
        last, now = 0, 0
        for num in nums:
            print(last, now, num)
            last, now = now, max(last + num, now)
        return now


if __name__ == '__main__':
    print(Solution().rob([2,7,9,3,1]))