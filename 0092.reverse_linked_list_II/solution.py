class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverse_between(self, head: ListNode, m: int, n: int) -> ListNode:
        if not head:
            return head
        ans = ListNode(None)
        ans.next = head
        p = ans
        for _ in range(m-1):
            p = p.next

        head = p.next
        for _ in range(m, n):
            temp = head.next
            head.next = temp.next
            temp.next = p.next
            p.next = temp
        return ans.next

