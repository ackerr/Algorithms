class Solution:
    def to_hex(self, num: int) -> str:
        if num == 0:
            return '0'
        prefix, ans = '0123456789abcdef', ''
        num = 2 ** 32 + num if num < 0 else num
        while num:
            item = num & 15  # num % 16
            ans = prefix[item] + ans
            num >>= 4  # num //= 4
        return ans
