class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def invert_tree(self, root: TreeNode) -> TreeNode:
        """ 对调left 与 right """
        if not root:
            return None
        root.left, root.right = root.right. root.left
        self.invert_tree(root.left)
        self.invert_tree(root.right)
        return root
