from typing import List


class Solution:
    def letter_combinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        m = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz",
        }
        ans = [""]
        for n in digits:
            ans = [pre + suf for pre in ans for suf in m.get(n)]
        return ans
