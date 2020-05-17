from utils import TreeNode


class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        return self.is_mirror(root, root)

    def is_mirror(self, left, right):
        if left is None and right is None:
            return True
        if not left or not right:
            return False
        return (
            left.val == right.val and self.is_mirror(left.left, right.right) and self.is_mirror(right.left, left.right)
        )
