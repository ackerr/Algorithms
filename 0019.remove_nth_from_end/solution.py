class ListNode:
    def __init__(self, v):
        self.v = v
        self.next = None


class Solution:
    def remove_nth_from_end(self, head: ListNode, n: int) -> ListNode:
        if not head:
            return
        ans = ListNode(0)
        ans.next = head
        fast, slow = ans, ans
        for _ in range(n + 1):
            fast = fast.next
        while fast.next:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next
        return ans.next
