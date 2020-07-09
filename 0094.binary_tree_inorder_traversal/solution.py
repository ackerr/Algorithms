from typing import List

import utils


class Solution:
    def inorder_traversal(self, root: utils.TreeNode) -> List[int]:
        def inorder(node, ans):
            if not node:
                return ans
            ans = inorder(node.left, ans)
            ans.append(node.val)
            ans = inorder(node.right, ans)
            return ans
        return inorder(root, [])

    def loop(self, root):
        stack, ans = [], []
        while root or stack:
            while root:
                stack.append(root)
                root = root.left
            root = stack.pop()
            ans.append(root.val)
            root = root.right
        return ans


