class Solution:
    def reverse_vowels(self, s: str) -> str:
        vowels = ["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"]
        left, right = 0, len(s) - 1
        words = list(s)
        while left < right:
            if words[left] in vowels and words[right] in vowels:
                words[left], words[right] = words[right], words[left]
                left, right = left + 1, right - 1
            if words[left] not in vowels:
                left += 1
            if words[right] not in vowels:
                right -= 1
        return "".join(words)
