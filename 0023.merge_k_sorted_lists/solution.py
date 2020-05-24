from typing import List

from utils import ListNode


class Solution:
    def violence(self, lists: List[ListNode]) -> ListNode:
        """ 暴力求解, 获取每个的链表所有元素，排序后组成链表"""
        values = []
        for link in lists:
            while link:
                values.append(link.val)
                link = link.next
        values.sort()
        ans = cur = ListNode(None)
        for value in values:
            cur.next = ListNode(value)
            cur = cur.next
        return ans.next

    def _merge_two_lists(self, l1: ListNode, l2: ListNode) -> ListNode:
        """ 合并两个链表 """
        ans = cur = ListNode(None)
        while l1 and l2:
            if l1.val < l2.val:
                cur.next = l1
                l1 = l1.next
            else:
                cur.next = l2
                l2 = l2.next
            cur = cur.next
        cur.next = l1 or l2
        return ans.next

    def merge_k_lists(self, lists: List[ListNode]) -> ListNode:
        """ 两两合并 """
        k = len(lists)
        if not lists or k == 0:
            return
        ans = lists[0]
        for i in range(1, k):
            ans = self._merge_two_lists(ans, lists[i])
        return ans

    def binary_merge_k_lists(self, lists: List[ListNode]) -> ListNode:
        """ 二分两两合并 """
        k = len(lists)
        if not lists or k == 0:
            return
        # 类似归并算法, 将排好序的链表合并为一个
        step = 1
        while step < k:
            for i in range(0, k - step, step * 2):
                lists[i] = self._merge_two_lists(lists[i], lists[i + step])
            step *= 2
        return lists[0]
