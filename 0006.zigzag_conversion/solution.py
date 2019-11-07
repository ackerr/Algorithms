class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows < 2:
            return s
        res = ["" for i in range(numRows)]
        index, flag = 0, -1

        for i in s:
            res[index] += i

            if index == 0 or index == numRows - 1:
                flag = -flag
            index += flag

        return "".join(res)


if __name__ == "__main__":
    assert Solution().convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR"
    assert Solution().convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI"

