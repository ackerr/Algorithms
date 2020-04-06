from typing import List


class Solution:
    def game_of_life(self, board: List[List[int]]):
        """
        Do not return anything, modify board in-place instead.
        """
        if not board:
            return board
        rows = len(board)
        cols = len(board[0])
        neighbors = [
            (-1, -1),
            (-1, 0),
            (-1, 1),
            (0, -1),
            (0, 1),
            (1, -1),
            (1, 0),
            (1, 1),
        ]
        for row in range(len(board)):
            for col in range(len(board[0])):
                count = 0
                for r, c in neighbors:
                    if (
                        0 <= r + row < rows
                        and 0 <= c + col < cols
                        and abs(board[r + row][c + col]) == 1
                    ):
                        count += 1

                # 原来活 现在死 记为 -1
                if (count < 2 or count > 3) and board[row][col] == 1:
                    board[row][col] = -1

                # 原来死 现在活 记为 2
                if count == 3 and board[row][col] == 0:
                    board[row][col] = 2

        for row in range(rows):
            for col in range(cols):
                board[row][col] = int(board[row][col] > 0)


if __name__ == "__main__":
    print(Solution().game_of_life([[0, 1, 0], [0, 0, 1], [1, 1, 1], [0, 0, 0]]))
