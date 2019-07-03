class Solution:
    def is_valid(self, s):
        mapping = {")": "(", "}": "{", "]": "["}
        stack = []
        for char in s:
            if char in "([{":
                stack.append(char)
            elif stack and stack[-1] == mapping[char]:
                stack.pop()
            else:
                return False
        return not stack
