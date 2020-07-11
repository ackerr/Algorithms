from typing import List

import utils


class Solution:
    def preorder_traversal(self, root: utils.TreeNode) -> List[int]:
        stack, ans = [root], []
        while stack:
            node = stack.pop()
            if node:
                ans.append(node.val)
                if node.right:
                    stack.append(node.right)
                if node.left:
                    stack.append(node.left)
        return ans
