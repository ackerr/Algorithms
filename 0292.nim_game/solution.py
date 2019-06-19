class Solution:
    def canWinNim(self, n: int) -> bool:
        """ 只要n不是4的倍数就一定赢 """
        return n % 4 != 0