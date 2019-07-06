class Solution:
    def is_perfect_square(self, num: int) -> bool:
        if num == 1:
            return True
        for i in range(num // 2 + 1):
            if i ** 2 == num:
                return True
            if i ** 2 > num:
                return False
        return False
