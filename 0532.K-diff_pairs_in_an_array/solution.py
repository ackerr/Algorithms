from typing import List


class Solution:
    def find_pairs(self, nums: List[int], k: int) -> int:
        if k < 0:
            return 0
        value, diff = set(), set()
        for i in nums:
            if i - k in value:
                diff.add(i - k)
            if i + k in value:
                diff.add(i)
            value.add(i)
        return len(diff)


if __name__ == "__main__":
    assert Solution().find_pairs([3, 1, 4, 1, 5], 2) == 2
