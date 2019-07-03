class Solution:
    def is_happy(self, n: int) -> bool:
        while n != 1 and n != 4:
            n = sum([int(i) ** 2 for i in str(n)])
        return n == 1
