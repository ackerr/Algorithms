class ListNode:
    def __init__(self, x)
        self.x = x
        self.next = None


class Solution:
    def swap_pairs(self, head: ListNode) -> ListNode:
        """ 每次向后移动两位, 交换前后节点值即可 """"
        p = head
        while p and p.next:
            p.val, p.next.val = p.next.val, p.val
            p = p.next.next
        return head
