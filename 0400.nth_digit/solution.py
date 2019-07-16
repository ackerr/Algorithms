class Solution:
    def find_nth_digit(self, n: int) -> int:
        for base in range(1, 11):
            base_value = 10 ** (base - 1)
            if n <= 9 * base_value * base:
                n -= 1
                return int(str(base_value + n // base)[n % base])
            n -= 9 * base_value * base

    def other_solution(self, n):
        base = 1
        while True:
            base_value = 10 ** (base - 1)
            if n <= 9 * base_value * base:
                n -= 1
                return int(str(base_value + n // base)[n % base])
            n -= 9 * base_value * base
            base += 1
