class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def oddEvenList(self, head: ListNode) -> ListNode:
        if not head:
            return head
        old, even = head, head.next
        even_head = head.next
        even = even
        while even and even.next:
            old.next = even.next
            even.next = even.next.next
            old = old.next
            even = even.next
        old.next = even_head
        return head
