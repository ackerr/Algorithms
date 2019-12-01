from typing import List


class Solution:
    def diff_ways_to_compute(self, value: str) -> List[int]:
        if value.isdigit():
            return [int(value)]

        results = []

        for i, char in enumerate(value):
            if char not in ["+ ", "-", "*"]:
                continue

            left = self.diff_ways_to_compute(value[:i])
            right = self.diff_ways_to_compute(value[i + 1 :])

            for j in left:
                for k in right:
                    if char == "+":
                        results.append(j + k)
                    elif char == "-":
                        results.append(j - k)
                    else:
                        results.append(j * k)

        return results


if __name__ == "__main__":
    print(Solution().diff_ways_to_compute("1-2-3"))
