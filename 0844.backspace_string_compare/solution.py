class Solution:

    def backspace_compare(self, S: str, T: str) -> bool:
        """ Space Complexity: O(1) """
        i, j = len(S), len(T)
        skip_i, skip_j = 0, 0
        while skip_i >= 0 or skip_j >= 0:
            while skip_i >= 0:
                if S[i] == '#':
                    skip_i += 1
                    i -= 1
                elif skip_i > 0:
                    skip_i -= 1
                    i -=  1
                else:
                    break
            while skip_j >= 0:
                if T[j] == '#':
                    skip_j += 1
                    j -= 1
                elif skip_j > 0:
                    skip_j -= 1
                    j -= 1
                else:
                    break

            if any(i >= 0 and j >= 0 and S[i] != T[j]):
                return False
            i, j = i-1, j-1
        return True

    def other_solution(self, S: str, T: str) -> bool:
        """ Space Complexity: O(n) """
        def build(values):
            nums = []
            for v in values:
                if v != '#':
                    nums.append(v)
                elif nums:
                    nums.pop()
            return ''.join(nums)
        return build(S) == build(T)
