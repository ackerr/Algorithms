from typing import List


class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        """ 因为len(nums) == 2n, 所以每次循环只需要添加nums[i]与nums[a+i] """
        ans = []
        for i in range(n):
            ans.append(nums[i])
            ans.append(nums[i + n])
        return ans

    def generic_template(self, nums, n):
        """ 如果给定数值长度不定 """
        ans = []
        length = len(nums)
        for i in range(n):
            index = i
            while index < length:
                ans.append(nums[index])
                index += n
        return ans
