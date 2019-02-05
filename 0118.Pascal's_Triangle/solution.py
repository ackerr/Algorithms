class Solution:

    def generate(self, num_rows):
        if num_rows == 0:
            return []
        ans = [1]
        result = [ans]
        for _ in range(num_rows-1):
            ans = [1]+[ans[i]+ans[i+1] for i in range(len(ans[:-1]))] + [1]
            result.append(ans)
        return result
