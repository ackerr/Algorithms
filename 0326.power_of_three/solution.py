class Solution:

    def is_power_of_three(self, n: int) -> bool:
        return n > 0 and 3 ** 19 % n == 0

    def loop_solution(self, n: int) -> bool:
        if n <= 0:
            return False
        while n != 1:
            if n % 3 != 0:
                return False
            n /= 3
        return True
