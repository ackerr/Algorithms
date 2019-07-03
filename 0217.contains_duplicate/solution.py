class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        """ 比较去重长度 跟 原始长度 即可 """
        return len(nums) != len(set(nums))
