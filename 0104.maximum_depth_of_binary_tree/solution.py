from utils import TreeNode


class Solution:
    def max_depth(self, root: TreeNode):
        if not root:
            return 0
        return max([self.max_depth(root.left), self.max_depth(root.right)]) + 1
