class Solution:

    def two_num(self, nums, target):
        new_dict = {}
        for index, value in enumerate(nums):
            if new_dict.get(target - value):
                return [new_dict.get(target - value), index]
            new_dict[value] = index
    
    def other_solution(self, nums, target):
        new_dict = {value: index for index, value in enumerate(nums)}
        for index, value in enumerate(nums):
            if new_dict.get(target - value) and index != new_dict[target - value]:
                return [index, new_dict[target - value]]

