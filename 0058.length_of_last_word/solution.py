class Solution:
    def length_of_last_word(self, s: str):
        for word in s.split(" ")[::-1]:
            if word:
                return len(word)
        return 0  # 处理 '', ' '


if __name__ == "__main__":
    print(Solution().length_of_last_word("  "))
