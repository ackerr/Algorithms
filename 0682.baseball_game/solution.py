from typing import List


class Solution:

    def cal_points(self, ops: List[str]) -> int:
        ans = []
        for op in ops:
            if op == 'C':
                ans.pop()
            elif op == 'D':
                ans.append(ans[-1] * 2)
            elif op == '+':
                ans.append(ans[-1]+ans[-2])
            else:
                ans.append(int(op))
        return sum(ans)


if __name__ == '__main__':
    print(Solution().cal_points(['5','2','C','D','+']))