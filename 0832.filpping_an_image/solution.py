from typing import List


class Solution:
    def flip_and_invert_image(self, a: List[List[int]]) -> List[List[int]]:
        for i in a:
            for j in range((len(i) + 1) // 2):
                i[j], i[~j] = i[~j] ^ 1, i[j] ^ 1
        return a

    def map_solution(self, a: List[List[int]]) -> List[List[int]]:
        return list(map(lambda i: [_ ^ 1 for _ in i][::-1], a))


if __name__ == '__main__':
    assert Solution().flip_and_invert_image([[1, 1, 0], [1, 0, 1], [0, 0, 0]]) == [[1, 0, 0], [0, 1, 0], [1, 1, 1, ]]
    assert Solution().map_solution([[1, 1, 0], [1, 0, 1], [0, 0, 0]]) == [[1, 0, 0], [0, 1, 0], [1, 1, 1, ]]
