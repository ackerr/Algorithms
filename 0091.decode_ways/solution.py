class Solution:
    def num_decodings(self, s: str) -> int:
        if s[0] == "0":  # 第一位为0无解
            return 0
        dp = [0] * (len(s) + 1)
        dp[0] = dp[1] = 1
        for i in range(1, len(s)):
            if s[i] == "0":  # 如果一个值为0，必须与前一个值组合，并且前面一个值必须为1或2才能组合
                if 0 < int(s[i - 1]) < 3:
                    dp[i + 1] += dp[i - 1]
                else:
                    return 0
            else:
                if 10 < int(s[i - 1:i + 1]) < 27:  #
                    dp[i + 1] += dp[i - 1] + dp[i]
                else:
                    dp[i + 1] += dp[i]
        return dp

    def optimize(self, s):
        """
        每次只需要前两位值，可以优化空间复杂度为O(1)
        """

        if int(s[0]) == 0:
            return 0
        pre = cur = 1
        for i in range(1, len(s)):
            if s[i] == "0":  # 如果一个值为0，必须与前一个值组合，并且前面一个值必须为1或2才能组合
                if 0 < int(s[i - 1]) < 3:
                    pre, cur = cur, pre
                else:
                    return 0
            else:
                if 10 < int(s[i - 1:i + 1]) < 27:  #
                    pre, cur = cur, pre + cur
                else:
                    pre, cur = cur, cur
        return cur
