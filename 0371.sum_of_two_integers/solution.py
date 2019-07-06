class Solution:
    def get_sum(self, a: int, b: int) -> int:
        """
        与  (a&b)<1  = 进位
        异或 a^b  = 无进位合

        直到进位等于0
        """
        if b == 0:
            return a
        return self.get_sum(a ^ b, (a & b) << 1)
