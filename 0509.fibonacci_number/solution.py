class Solution:
    def fib_loop(self, n: int) -> int:
        a, b = 0, 1
        for _ in range(n):
            a, b = b, a + b
        return a

    def fib_recursion(self, n: int) -> int:
        if n < 2:
            return n
        return self.fib_recursion(n - 1) + self.fib_recursion(n - 2)


if __name__ == "__main__":
    assert Solution().fib_loop(10) == 55
    assert Solution().fib_recursion(10) == 55
