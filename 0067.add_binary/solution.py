class Solution:

    def add_binary(self, a, b):
        # 'b' to Binary; 'o' to Octal;
        return format(int(a, 2) + int(b, 2), 'b')

    def other_solution(self, a, b):
        # int(num: str, 2)  str(binary) -> int
        # bin(num: int)     int -> str(binary)
        return bin(int(a, 2) + int(b, 2))[2:]
