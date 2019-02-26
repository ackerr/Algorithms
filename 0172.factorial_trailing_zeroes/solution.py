class Solution:

    def trailing_zeroes(self, n: int) -> int:
        """
        因式分解会发现只有当出现 2x5 的时候才会末尾有 0 , 又因为出现2的概率比5高多了，所以只需要计算出现5的个数即可
        """
        total = 0
        while n:
            n //= 5
            total += n
        return total

    def other_solution(self, n):
        """ n//5 + n//25 + n//125 + ... """
        total = 0
        while 5**total <= n:
            total += 1
        return sum([n//5**i for i in range(1, total)])
