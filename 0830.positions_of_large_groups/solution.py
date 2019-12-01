from typing import List


class Solution:
    def large_group_positions(self, s: str) -> List[List[int]]:
        ans, left, right = [], 0, 0
        for i in range(len(s)):
            if s[i] == s[right]:
                right = i
            else:
                left = right = i
            if (
                right - left > 1 and s[right] != s[right + 1 : right + 2]
            ):  # tip: ignore out of index error
                ans.append([left, right])
        return ans


if __name__ == "__main__":
    assert Solution().large_group_positions("abbxxxxzzy") == [[3, 6]]
    assert Solution().large_group_positions("aaa") == [[0, 2]]
