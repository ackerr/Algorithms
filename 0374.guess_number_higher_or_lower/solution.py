# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0


def guess(num):
    pass


class Solution:
    def guess_number(self, n: int) -> int:
        left, right = 1, n + 1
        while left < right:
            mid = (right + left) >> 1
            if guess(mid) == 1:
                left = mid + 1
            elif guess(mid) == -1:
                right = mid - 1
            else:
                return mid
        return left
