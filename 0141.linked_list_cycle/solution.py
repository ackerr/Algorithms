# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def has_cycle(self, head: ListNode) -> bool:
        """ 如果是环形链表，一快一慢的双指针最终会相遇 """
        if not head or not head.next:  # 防止head为空或只为一个值
            return False
        fast = head.next  # 快指针
        slow = head  # 慢指针
        while fast != slow:
            if not fast or not fast.next:
                return False
            fast = fast.next.next
            slow = slow.next
        return True
