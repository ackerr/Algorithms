class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def middle_node(self, head: ListNode) -> ListNode:
        fast, slow = head, head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        return slow
