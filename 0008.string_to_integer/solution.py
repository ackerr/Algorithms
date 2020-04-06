import re


class Solution:
    def my_atoi(self, s: str) -> int:
        """ int(*[]) == 0 """
        return max(
            min(int(*re.findall(r"^[\-\+]?\d+", s.lstrip())), (2 << 30) - 1), -2 << 30
        )


if __name__ == "__main__":
    print(Solution().my_atoi("    -41"))
    print(Solution().my_atoi("+    -"))
    print(Solution().my_atoi("  hello world  -41 hi"))
    print(Solution().my_atoi("words and 987"))
    print(Solution().my_atoi("2147483648"))
