from utils import ListNode


class Solution:
    def insertion_sort_list(self, head: ListNode) -> ListNode:
        """ 因为单向链表，不能像数组一样从后向前比较，而是从前往后直接找到对于的位置插入 """
        if not head or not head.next:  # empty or one element
            return head
        ans = ListNode(0)
        ans.next = head
        while head and head.next:
            if head.val <= head.next.val:  # greater than prev
                head = head.next
                continue
            pre = ans
            while pre.next.val < head.next.val:  # insert index
                pre = pre.next
            temp = head.next  # move element
            head.next = temp.next
            temp.next = pre.next
            pre.next = temp
        return ans.next
