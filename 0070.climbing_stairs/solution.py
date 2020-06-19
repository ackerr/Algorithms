class Solution:
    """
    n=1, res=1
    n=2, res=2
    n=3, res=3
    n=4, res=5
    n=5, res=8
    ...
    可以看出就是斐波那契的变种
    """

    def climb_stairs(self, n):
        if n < 3:
            return n
        return self.climb_stairs(n - 1) + self.climb_stairs(n - 2)

    def loop_solution(self, n):
        a, b = 1, 1
        for _ in range(n):
            a, b = b, a + b
        return a


if __name__ == "__main__":
    for i in range(1, 10):
        print(Solution().climb_stairs(i))
        print(Solution().loop_solution(i))
