class Solution:
    def is_power_of_two(self, n: int) -> bool:
        “”“
        二进制下 
            - 假设n 为2的幂
            - n，首位为1,其余位为0
            - n-1 ， 全部位为1，长度减1
            - n & n-1 ,则为0
        “”“
        return n > 0 and (n & n-1) == 0