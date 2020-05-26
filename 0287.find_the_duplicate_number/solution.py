from typing import List


class Solution:
    def find_duplicate(self, nums: List[int]) -> int:
        """ 类型环形链表II """
        slow = fast = 0
        fast = nums[nums[fast]]
        slow = nums[slow]
        while nums[slow] != nums[fast]:
            fast = nums[nums[fast]]
            slow = nums[slow]
        fast = 0
        while fast != slow:
            slow = nums[slow]
            fast = nums[fast]
        return fast
