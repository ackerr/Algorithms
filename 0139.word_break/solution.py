from typing import List


class Solution:
    def word_break(self, s: str, word_dict: List[str]) -> bool:
        length = len(s)
        dp = [False] * (length + 1)
        dp[0] = True
        word_dict = set(word_dict)
        for i in range(1, length+1):
            for j in range(i):
                if dp[j] and s[j:i] in word_dict:
                    dp[i] = True
                    break
        return dp[-1]
