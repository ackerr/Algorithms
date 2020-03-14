class Solution:
    def longest_increasing_subsequence(self, nums):
        length = len(nums)
        if length < 2:
            return length
        dp = [1] * length
        for i in range(length):
            for j in range(0, i):
                if nums[i] > nums[j]:
                    dp[i] = max(dp[i], dp[j] + 1)
        return max(dp)


if __name__ == "__main__":
    print(Solution().longest_increasing_subsequence([10, 9, 2, 5, 3, 7, 101, 18]))
