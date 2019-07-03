class Solution:
    def is_isomorphic(self, s: str, t: str) -> bool:
        """ 根据值在两个字符串中的索引是否相同 """
        for i, v in enumerate(s):
            if s.find(v) != t.find(t[i]):
                return False
        return True

    def other_solution(self, s, t):
        """ 将s中的字符对应的t中的字符存入字典，如果不符合条件则为False """
        d = {}
        for i, v in enumerate(s):
            if d.get(v):
                if d[v] != t[i]:
                    return False
            elif t[i] in d.values():
                return False
            d[v] = t[i]
        return True


if __name__ == "__main__":
    print(Solution().test("abc", "bac"))
