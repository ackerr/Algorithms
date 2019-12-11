class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        sign = (dividend < 0) == (divisor < 0)  # 判断正负
        a, b, res = abs(dividend), abs(divisor), 0
        while a >= b:
            x = 0
            while a >= b << (x + 1):  # 每次将b左移一位(乘2)，比较大小, 缩小范围
                x += 1
            res += 1 << x
            a -= b << x
        return min(res if sign else -res, 2**31-1)  # 防一手溢出
