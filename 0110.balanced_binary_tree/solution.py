from utils import TreeNode


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
