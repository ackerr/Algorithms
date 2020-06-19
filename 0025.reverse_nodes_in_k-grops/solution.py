import utils


class Solution:
    def reverse(self, head: utils.ListNode, tail: utils.ListNode):
        """ 反转head到tail之间的节点 """
        prev = tail.next
        p = head
        while prev != tail:
            temp = p.next
            p.next = prev
            prev = p
            p = temp
        return tail, head

    def reverse_k_group(self, head: utils.ListNode, k: int) -> utils.ListNode:
        ans = utils.ListNode(0)
        ans.next = head
        cur = ans

        while head:
            tail = cur
            # 剩余部分长度是否大于等于k
            for i in range(k):
                tail = tail.next
                if not tail:
                    return ans.next
            temp = tail.next
            head, tail = self.reverse(head, tail)
            # 把子链表重新接回原链表
            cur.next = head
            tail.next = temp
            cur = tail
            head = tail.next
        return ans.next
