class Solution:

    def strStr(self, haystack: str, needle: str) -> int:
        if needle == '':
            return 0
        if haystack.split(needle) != haystack:
            return len(haystack.split(needle))
        return -1

    def use_find_solution(self, haystack: str, needle: str) -> int:
        return haystack.find(needle)
