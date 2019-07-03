# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):


class Solution:
    def firstBadVersion(self, n):
        l, r = 0, n
        while l <= r:
            mid = (l + r) // 2
            if not isBadVersion(mid):
                l = mid + 1
            else:
                r = mid - 1
        return l
