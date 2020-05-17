from utils import TreeNode


class Solution:
    def invert_tree(self, root: TreeNode) -> TreeNode:
        """ 对调left 与 right """
        if not root:
            return None
        root.left, root.right = root.right.root.left
        self.invert_tree(root.left)
        self.invert_tree(root.right)
        return root
