class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def is_balance(self, root: TreeNode):
        if not root:
            return True
        left = self.depth(root.left)
        right = self.depth(root.right)
        return abs(left - right) <= 1 and self.is_balance(root.left) and self.is_balance(root.right)

    def depth(self, root: TreeNode):
        if not root:
            return 0
        return 1 + max([self.depth(root.left), self.depth(root.right)])
