# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def is_same_tree(self, p: TreeNode, q: TreeNode) -> bool:
        if p is None and q is None:
            return True
        if not p or not q:
            return False
        return p.val == q.val and self.is_same_tree(p.left, q.left) and self.is_same_tree(p.right, q.right)