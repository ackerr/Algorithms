from utils import ListNode


class Solution:
    def add_two_numbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        res = ans = ListNode(0)
        pon = 0

        while pon or l1 or l2:
            i = (l1.val if l1 else 0) + (l2.val if l2 else 0) + pon
            pon = i // 10
            ans.next = ListNode(i % 10)
            ans = ans.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None

        return res.next
