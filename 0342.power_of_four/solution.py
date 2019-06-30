class Solution:
    """ 与 0326.power_of_three 类似"""

    def is_power_of_four(self, num: int) -> bool:
        """ 2的倍数且前一个数是3的倍数 """
        return num & (num - 1) == 0 and num % 3 == 1

    def loop_solution(self, num: int) -> bool:
        if num <= 0:
            return False
        while num != 1:
            if num % 4 != 0:
                return False
            num /= 4
        return True
