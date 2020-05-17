from utils import ListNode


class Solution:
    def reverse_between(self, head: ListNode, m: int, n: int) -> ListNode:
        if not head:
            return head
        ans = ListNode(None)
        ans.next = head
        p = ans
        for _ in range(m - 1):
            p = p.next

        head = p.next
        for _ in range(m, n):
            temp = head.next
            head.next = temp.next
            temp.next = p.next
            p.next = temp
        return ans.next
