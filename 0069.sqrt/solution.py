class Solution:
    def my_sqrt(self, x):
        return int(x ** 0.5)

    def binary_search(self, x):
        if x < 2:
            return x
        left, right = 1, x
        while left < right:
            mid = (left + right) // 2
            if mid ** 2 > x:
                right = mid
            else:
                left = mid + 1
        return left - 1


if __name__ == "__main__":
    print(Solution().binary_search(10))
