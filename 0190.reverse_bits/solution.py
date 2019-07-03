class Solution:
    def reverse_bits(self, n):
        """
        2 << 1 == 2 * 2
        2 >> 1 == 2 / 2
        1 & 1 == 1
        1 & 0 == 0
        0 & 0 == 0

        """
        ans = 0
        for i in range(32):
            ans = (ans << 1) + (n & 1)
            n >>= 1
        return ans
