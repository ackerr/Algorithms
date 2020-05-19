class Solution:
    def valid_palindrome(self, s: str) -> bool:
        if not s:
            return True
        left, right = 0, len(s) - 1
        while left < right:
            if s[left] == s[right]:
                left += 1
                right -= 1
            else:  # 出现不一致时，去掉左边或右边元素。判断是否是回文
                a = s[left + 1 : right + 1]
                b = s[left:right]
                return a == a[::-1] or b == b[::-1]
        return True
