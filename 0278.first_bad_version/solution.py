# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool


def isBadVersion(version):
    pass


class Solution:
    def firstBadVersion(self, n):
        left, right = 0, n
        while left <= right:
            mid = (left + right) // 2
            if not isBadVersion(mid):
                left = mid + 1
            else:
                right = mid - 1
        return left
