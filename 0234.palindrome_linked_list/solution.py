from utils import ListNode


class Solution:
    # def is_palindrome(self, head: ListNode) -> bool:
    #     """ 时间复杂度：O(n)  空间复杂度: O(n)"""
    #     if not head:
    #         return True
    #     ans = []
    #     while head:
    #         ans.append(head.val)
    #         head = head.next
    #     return ans == ans[::-1]

    def is_palindrome(self, head: ListNode) -> bool:
        """ 通过快慢指针，找到后半截链表，反转后半截链表与原链表比较值 """
        if not head:
            return True

        # 快慢指针获取后半截链表，需要注意奇偶个数问题
        fast, slow, pre = head, head, None
        while fast:
            fast = fast.next.next if fast.next else fast.next
            slow = slow.next

        # 反转链表
        while slow:
            temp = slow.next
            slow.next = pre
            pre = slow
            slow = temp

        # 比较值
        while pre:
            if head.val != pre.val:
                return False
            head = head.next
            pre = pre.next
        return True
