class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        """ 
        只要有一个重复值的索引差 小于等于 k 即可 
        通过dict存值，出现重复指时，进行比较索引值，如果大于则覆盖原有索引
        """
        storage = {}
        for i, value in enumerate(nums):
            if value in storage:
                if (i - storage[value]) <= k:
                    return True
            storage[value] = i
        return False