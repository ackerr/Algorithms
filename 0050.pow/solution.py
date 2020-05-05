class Solution:
    def my_pow(self, x: float, n: int) -> float:
        if n == 0:
            return 1
        if n < 0:
            return 1 / self.my_pow(x, -n)
        if n % 2:
            return x * self.my_pow(x, n - 1)
        return self.my_pow(x * x, n // 2)
