class Solution:
    
    def reverse(self, x):
        # 2**31 == 2147483648 == (2<<30)
        ans = int(('-' if x < 0 else '' ) + str(abs(x))[::-1])
        return 0 if ans < (-2<<30) or ans > (2<<30) else ans

