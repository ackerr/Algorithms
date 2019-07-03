class Solution:
    def roman_to_int(self, s):
        chart = {"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
        total_value = sum([chart[value] for value in s])
        for index, value in enumerate(s[:-1]):
            if chart[value] < chart.get(s[index + 1]):
                total_value -= chart.get(s[index]) * 2
        return total_value

    # other solution
    def solution_by_if(self, s):
        chart = {"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
        total_value = sum([chart[value] for value in s])
        if "CM" in s:
            total_value -= 200
        if "CD" in s:
            total_value -= 200
        if "XC" in s:
            total_value -= 20
        if "XL" in s:
            total_value -= 20
        if "IX" in s:
            total_value -= 2
        if "IV" in s:
            total_value -= 2
        return total_value

    def solution_by_replace(self, s):
        chart = {
            "M": 1000,
            "D": 500,
            "C": 100,
            "L": 50,
            "X": 10,
            "V": 5,
            "I": 1,
            "a": 900,
            "b": 400,
            "c": 90,
            "d": 40,
            "e": 9,
            "f": 4,
        }
        s = (
            s.replace("CM", "a")
            .replace("CD", "b")
            .replace("XC", "c")
            .replace("XL", "d")
            .replace("IX", "e")
            .replace("IV", "f")
        )
        return sum([chart[value] for value in s])


if __name__ == "__main__":
    s = Solution()
    print(s.roman_to_int("MCMIVIII"))
    print(s.solution_by_if("MCMIVIII"))
    print(s.solution_by_replace("MCMIVIII"))
