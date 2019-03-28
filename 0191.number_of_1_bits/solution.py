class Solution:

    def hamming_weight(self, n):
        count = 0
        for i in range(32):
            if (n & 1) == 1:
                count += 1
            n >>= 1
        return count

    def other_solution(self, n):
        """ use build in function """
        return bin(n).count('1')
