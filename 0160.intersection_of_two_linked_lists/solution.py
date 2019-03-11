class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:

    def get_intersection_node(self, headA: ListNode, headB: ListNode) -> ListNode:
        a, b = headA, headB
        while a != b:
            a = a.next if a else headB
            b = b.next if b else headA
        return a
