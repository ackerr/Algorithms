class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def rotate_right(self, head: ListNode, k: int) -> ListNode:
        """ k%length 为 末尾需要移动到头部的节点数 """
        if not head or not head.next or k <= 0:
            return head
        count = 1
        p = f = head
        while p.next:
            count += 1
            p = p.next
        if k % count == 0:
            return head
        p.next = head  # 先将原尾节点指向原头结点
        for _ in range(count - k % count - 1):
            f = f.next
        ans = f.next  # 取新头结点
        f.next = None  # 新尾节点
        return ans
