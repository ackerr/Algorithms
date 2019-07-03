class Solution:
    def is_palindrome(self, s):
        # str.isalnum  判断字符串是否由字母或数字组成
        s = "".join(i for i in s if i.isalnum()).lower()
        return s == s[::-1]

    def re_solution(self, s):
        # re.sub()  替换正则匹配出的内容
        import re

        s = re.sub("[^0-9a-zA-Z]", "", s).lower()
        return s == s[::-1]


if __name__ == "__main__":
    print(Solution().is_palindrome("A man, a plan, a canal: Panama"))
    print(Solution().re_solution("A man, a plan, a canal: Panama"))
