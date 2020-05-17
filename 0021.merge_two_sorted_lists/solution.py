from utils import ListNode


class Solution:
    """ 用上方的类模拟链表 """

    def merge_two_lists(self, l1: ListNode, l2: ListNode) -> ListNode:
        cur = ans = ListNode(0)  # 定义两个空链表，一个用于过程，一个用于结尾
        while l1 and l2:
            if l1.val > l2.val:  # 接上小的值
                cur.next = l2
                l2 = l2.next
            else:
                cur.next = l1
                l1 = l1.next
            cur = cur.next
        cur.next = l1 or l2  # 填补多出的值
        return ans.next  # 去掉初始化的0
