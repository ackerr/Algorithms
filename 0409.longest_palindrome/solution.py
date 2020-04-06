import collections


class Solution:
    def longestPalindrome(self, s: str) -> int:
        count = collections.Counter(s)
        res = 0
        for k in count:
            res += count[k] // 2 * 2
        return res if res == len(s) else res + 1
