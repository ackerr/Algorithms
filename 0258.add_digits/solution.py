class Solution:
    def add_digits(self, num: int) -> int:
        """
        original_value = 100*a + 10*b + 1*c
        value = a + b + c
        original_value - value = 99 * a + 9 * b
        (original_value - value) % 9 == 0
        """
        if num < 10:
            return num
        num = num % 9
        return 9 if num == 0 else num

    def other_solution(self, num):
        while len(str(num)) > 1:
            num = sum(map(int, str(num)))
        return num
