class Solution:

    def plus_one(self, digits: list):
        # 最后一位不是9的标记位
        length = len(digits) - 1
        while digits[length] == 9:
            digits[length] = 0
            length -= 1
        if length < 0:
            digits = [1] + digits
        else:
            digits[length] += 1
        return digits

    def other_solution(self, digits: list):
        # 将数组合并+1，再拆分，注意输出类型就ok
        # 时间复杂度o(n) 不太理想
        temp = ''.join(map(lambda i: str(i), digits))
        return [int(n) for n in str(int(temp) + 1)]


if __name__ == '__main__':
    print(Solution().plus_one([9]))
    print(Solution().other_solution([9]))
