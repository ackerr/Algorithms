class Solution:

    def is_ugly(self, num: int) -> bool:
        if num == 0:
            return False
        if num == 1:
            return True
        if num % 2 == 0:
            return self.isUgly(num/2)
        if num % 3 == 0:
            return self.isUgly(num/3)
        if num % 5 == 0:
            return self.isUgly(num/5)
        return False