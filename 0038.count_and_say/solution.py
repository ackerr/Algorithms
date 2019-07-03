import re


class Solution:
    def count_and_say(self, n):
        data = "1"
        for _ in range(n - 1):
            data = re.sub(r"(.)\1*", lambda m: str(len(m.group(0)) + m.group(1)), n)
        return data


if __name__ == "__main__":
    print(re.findall(r"(.)\1*", "aaaabcd"))  # ['a', 'b', 'c', 'd']
