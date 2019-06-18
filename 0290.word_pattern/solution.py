class Solution:
    def word_pattern(self, pattern: str, str: str) -> bool:
        words = str.split(' ')
        return len(pattern) == len(words) and len(set(pattern)) == len(set(words)) == len(set(zip(pattern, words)))