from typing import List


class Solution:
    def group_anagrams(self, strs: List[str]) -> List[List[str]]:
        ans = {}
        for i in strs:
            j = str(sorted(i))
            if j not in ans.keys():
                ans[j] = [i]
            else:
                ans[j].append(i)
        return [i for i in ans.values()]


if __name__ == "__main__":
    assert Solution().group_anagrams(["abc", "cba", "baa", "aba"]) == [["abc", "cba"], ["baa", "aba"]]
