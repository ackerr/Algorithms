from typing import List


class Solution:
    def generate_parenthesis(self, n: int) -> List[str]:
        def backtrack(s="", left=0, right=0):
            if len(s) == 2 * n:
                ans.append(s)
                return
            if left < n:
                backtrack(s + "(", left + 1, right)
            if right < left:
                backtrack(s + ")", left, right + 1)

        ans = []
        backtrack()
        return ans


if __name__ == "__main__":
    assert len(Solution().generate_parenthesis(2)) == 2
    assert len(Solution().generate_parenthesis(3)) == 5
