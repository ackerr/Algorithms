from utils import ListNode


class Solution:
    def delete_duplicates(self, head: ListNode) -> ListNode:
        if not head or not head.next:
            return head
        node = ListNode(None)
        node.next = head
        target = node
        while target.next:
            left = right = target.next
            while right.next and right.next.val == left.val:  # duplicate
                right = right.next
            if left == right:  # no move
                target = target.next
            else:
                target.next = right.next  # remove duplicate node
        return node.next
