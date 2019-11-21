class Solution:
    def search_matrix(self, matrix, target):
        """
        从左下角开始
        1. 如果左下角的值大于target, 则此行都大于
        2. 如果左下角的值小于target, 则向右移动一位
        3. 重复操作直到找到一个值或到达边界
        """
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return False
        m, n = len(matrix) - 1, 0
        while m >= 0 and n < len(matrix[0]):
            if matrix[m][n] == target:
                return True
            elif matrix[m][n] > target:
                m -= 1
            else:
                n += 1
        return False


if __name__ == '__main__':
    matrix = [[1, 5, 9, 13, 17], [2, 6, 10, 14, 18], [3, 7, 11, 15, 19], [4, 8, 12, 16, 20]]
    assert Solution().search_matrix(matrix, 12)
    assert not Solution().search_matrix(matrix, 23)
