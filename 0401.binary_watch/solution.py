from typing import List


class Solution:
    def read_binary_watch(self, num: int) -> List[str]:
        ans = []
        for i in range(12):
            for j in range(60):
                if bin(i).count("1") + bin(j).count("1") == num:
                    ans.append("%d:%02d" % (i, j))
        return ans
