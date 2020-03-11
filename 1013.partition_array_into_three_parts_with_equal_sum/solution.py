from typing import List


class Solution:
    def canThreePartsEqualSum(self, A: List[int]) -> bool:
        s = sum(A)
        if s % 3 != 0:
            return False
        target = s // 3
        left, right, cur = None, None, 0
        for i, n in enumerate(A):
            cur += n
            if cur == target:
                if left is None:
                    left = i
                else:
                    right = i
                    break
                cur = 0
        return bool(right and right != len(A) - 1 and sum(A[right + 1 :]) == target)


if __name__ == "__main__":
    assert not Solution().canThreePartsEqualSum([24, -4, -5, -12, 5, 16, -12, 22, 2]))
    assert Solution().canThreePartsEqualSum([12, -4, 16, -5, 9, -3, 3, 8, 0]))
