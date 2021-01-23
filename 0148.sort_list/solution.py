from utils import ListNode


class Solution:
    def sort_list(self, head: ListNode) -> ListNode:
        """
        时间复杂度：O(nlogn)
        空间复杂度：O(1)
        """

        def split(head: ListNode, tail: ListNode) -> ListNode:
            """ 递归拆分链表 """
            if not head:
                return head
            if head.next == tail:
                head.next = None
                return head
            slow = fast = head
            while fast != tail:
                slow = slow.next
                fast = fast.next
                if fast != tail:
                    fast = fast.next
            mid = slow
            return merge_sort(split(head, mid), split(mid, tail))

        def merge_sort(head1: ListNode, head2: ListNode) -> ListNode:
            """ 合并两个有序链表 """
            ans = ListNode(0)
            temp, temp1, temp2 = ans, head1, head2
            while temp1 and temp2:
                if temp1.val <= temp2.val:
                    temp.next = temp1
                    temp1 = temp1.next
                else:
                    temp.next = temp2
                    temp2 = temp2.next
                temp = temp.next
            temp.next = temp1 or temp2
            return ans.next

        return split(head, None)

    def other_solution(self, head: ListNode) -> ListNode:
        """
        时间复杂度：O(nlogn)
        空间复杂度：O(n)
        """
        if not head:
            return head
        v = []
        cur = head
        while cur:
            v.append(cur.val)
            cur = cur.next
        ans = head
        v.sort()
        for i in v:
            ans.val = i
            ans = ans.next
        return head
