# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        return self.is_mirror(root, root)
    
    def is_mirror(self, left, right):
        if left is None and right is None:
            return True
        if not left or not right:
            return False
        return left.val == right.val and self.is_mirror(left.left, right.right) and self.is_mirror(right.left, left.right)