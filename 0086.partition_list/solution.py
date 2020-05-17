from utils import ListNode


class Solution:
    def partition(self, head: ListNode, x: int) -> ListNode:
        if not head:
            return head
        ans = low = ListNode(None)
        target = high = ListNode(None)
        while head:
            if head.val < x:
                low.next = head
                low = low.next
            else:
                high.next = head
                high = high.next
            head = head.next
        high.next = None
        low.next = target.next
        return ans.next
