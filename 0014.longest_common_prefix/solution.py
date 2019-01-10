class Solution:

    def longest_common_prefix(self, strs):
        prefix = ''
        if len(strs) == 0:
            return prefix

        for i in range(len(strs[0])):
            char = strs[0][i]
            value = sum([s[i] == char for s in strs if i < len(s)])
            if value == len(strs):
                prefix += char
            else:
                break
        return prefix


if __name__ == '__main__':
    s = Solution()
    print(s.longest_common_prefix(['s']))
    print(s.longest_common_prefix(['s', 'as']))
