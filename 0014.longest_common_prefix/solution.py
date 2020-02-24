class Solution:
    def longest_common_prefix(self, strs):
        if len(strs) == 0:
            return ""

        m, n = min(strs), max(strs)
        for i, v in enumerate(m):
            if v != n[i]:
                return m[:i]
        return m


if __name__ == "__main__":
    s = Solution()
    print(s.longest_common_prefix(["s", "as"]))
    print(s.longest_common_prefix(["flower", "flow", "flight"]))
