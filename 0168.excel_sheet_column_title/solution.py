class Solution:
    def convert_to_title(self, n: int) -> str:
        """
        1 -> 'A'
        27 -> 'AA'
        703 -> 'AAA'
        """
        ans = []
        while n > 0:
            n, mod = divmod(n, 26)
            if mod == 0:
                mod = 26
                n -= 1
            ans.append(chr(64 + mod))
        return "".join(ans[::-1])
