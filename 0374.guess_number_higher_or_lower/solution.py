# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0


def guess(num):
    pass


class Solution:
    def guess_number(self, n: int) -> int:
        l, r = 1, n + 1
        while l < r:
            mid = (r + l) >> 1
            if guess(mid) == 1:
                l = mid + 1
            elif guess(mid) == -1:
                r = mid - 1
            else:
                return mid
        return l
