from utils import TreeNode


class Solution:
    def is_subtree(self, s: TreeNode, t: TreeNode) -> bool:
        """ 判断本身/左子树/右子树是不是相同"""
        if not s and not t:
            return True
        elif not s or not t:
            return False
        return self.is_same_tree(s, t) or self.is_same_tree(s.left, t) or self.is_same_tree(s.right, t)

    def is_same_tree(self, m: TreeNode, n: TreeNode) -> bool:
        """ 判断两树是不是相同 (like 100.same_tree)"""
        if not m and not n:
            return True
        elif not m or not n:
            return False
        return m.val == n.val and self.is_same_tree(m.left, n.left) and self.is_same_tree(m.right, n.right)
