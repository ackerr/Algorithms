class Solution:
    def find_the_difference(self, s: str, t: str) -> str:
        for c in t:
            if c not in s:
                return c
            else:
                s = s.replace(c, '', 1)
