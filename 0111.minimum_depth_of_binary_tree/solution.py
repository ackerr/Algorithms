class Solution:
    def min_depth(self, root):
        if not root:
            return 0
        left = self.min_depth(root.left)
        right = self.min_depth(root.right)
        if not left or not right:
            return left + right + 1
        return min([left, right]) + 1
