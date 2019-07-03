class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverse_list(self, head: ListNode) -> ListNode:
        cur, ans = head, None
        while cur:
            temp = cur.next
            cur.next = ans
            ans = cur
            cur = temp
        return ans

    def recursion_solution(self, head: ListNode, ans=None) -> ListNode:
        if not head:
            return ans
        temp = head.next
        head.next = ans
        ans = head
        head = temp
        return self.recursion_solution(head, ans)
