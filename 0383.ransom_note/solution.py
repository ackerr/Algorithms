class Solution(object):
    def can_construct(self, ransom_note: str, magazine: str) -> bool:
        # from collections import Counter
        # n, m = Counter(ransom_note), Counter(magazine)
        n, m = {}, {}
        for i in ransom_note:
            if i in n:
                n[i] += 1
            else:
                n[i] = 1
        for i in magazine:
            if i in m:
                m[i] += 1
            else:
                m[i] = 1

        for k, v in n.items():
            if k not in m or v > m[k]:
                return False
        return True

    def other_solution(self, ransom_note: str, magazine: str) -> bool:
        for i in ransom_note:
            result = magazine.replace(i, "", 1)
            if len(result) == len(magazine):
                return False
            magazine = result
        return True
