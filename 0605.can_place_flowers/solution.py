from typing import List


class Solution:
    def can_place_flowers(self, flowerbed: List[int], n: int) -> bool:
        ans = 0
        temp = [0] + flowerbed + [0]
        for i in range(1, len(temp) - 1):
            if temp[i - 1] == temp[i + 1] == temp[i] == 0:
                ans += 1
                temp[i] = 1
        return ans >= n


if __name__ == '__main__':
    assert not Solution().can_place_flowers([1, 0, 0, 1, 0, 0, 1], 1)
    assert Solution().can_place_flowers([0, 0, 1, 0, 0], 2)
