class Solution:
    def length_of_longest_substring(self, s: str) -> int:
        count = res = 0
        sign = {}
        for i in range(len(s)):
            if s[i] in sign:  # duplicate
                count = max(sign[s[i]], count)
            sign[s[i]] = i + 1  # update index
            res = max(res, i - count + 1)  # update max value
        return res


if __name__ == "__main__":
    assert Solution().length_of_longest_substring("abba") == 2
