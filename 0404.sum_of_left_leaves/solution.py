class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def sum_of_left_leaves(self, root: TreeNode) -> int:
        if not root:
            return 0
        if root.left and not root.left.left and not root.left.right:
            return root.left.val + self.sum_of_left_leaves(root.left)
        else:
            return self.sum_of_left_leaves(root.left) + self.sum_of_left_leaves(root.right)