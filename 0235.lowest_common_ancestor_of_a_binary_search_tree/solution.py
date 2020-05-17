from utils import TreeNode


class Solution:
    def lowest_common_ancestor(self, root: "TreeNode", p: "TreeNode", q: "TreeNode") -> "TreeNode":
        if root.val > p.val and root.val > q.val:
            return self.lowest_common_ancestor(root.left, p, q)
        if root.val < p.val and root.val < q.val:
            return self.lowest_common_ancestor(root.right, p, q)
        return root
