class Solution:
    def two_sum(self, numbers, target):
        flag = {}
        for index, number in enumerate(numbers):
            sub = target - number
            if sub in flag.keys():
                return flag[sub] + 1, index + 1
            else:
                flag[number] = index

    def loop_solution(self, numbers, target):
        left, right = 0, len(numbers) - 1
        while numbers[left] + numbers[right] != target:
            if numbers[left] + numbers[right] < target:
                left += 1
            else:
                right -= 1
        return left + 1, right + 1
