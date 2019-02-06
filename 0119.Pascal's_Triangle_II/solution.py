class Solution:

    def get_row(self, row_index):
        ans = [1]
        for _ in range(row_index):
            ans = [1] + [ans[i] + ans[i+1] for i in range(len(ans) - 1)] + [1]
        return ans
