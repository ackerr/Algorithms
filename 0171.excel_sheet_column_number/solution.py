class Solution:
    def title_to_number(self, s: str) -> int:
        """ 26进制 """
        ans = 0
        for value in s:
            ans = ans * 26 + ord(value) - 64
        return ans
