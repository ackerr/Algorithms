class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def detect_cycle(self, head: ListNode):
        fast, slow = head, head
        while fast and fast.next:
            fast, slow = fast.next.next, slow.next
            if fast == slow:
                break
        if not(fast and fast.next):
            return
        while head != slow:
            head, slow = head.next, slow.next
        return head

    def hash_solution(self, head: ListNode):
        if not head:
            return None
        node = set()
        while head:
            if head in node:
                return head
            else:
                node.add(head)
            head = head.next
