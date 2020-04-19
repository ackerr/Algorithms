class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def removeElements(self, head: ListNode, val: int) -> ListNode:
        if not head:
            return head
        cur = head
        while cur:
            if cur.next and cur.next.val == val:
                cur.next = cur.next.next
            else:
                cur = cur.next
        return head.next if head.val == val else head  # verify the case of linked-list length is one
