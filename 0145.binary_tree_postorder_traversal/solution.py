from typing import List

import utils


class Solution:
    def postorder_traversal(self, root: utils.TreeNode) -> List[int]:
        def dfs(node, ans):
            if not node:
                return ans
            dfs(node.left, ans)
            dfs(node.right, ans)
            ans.append(node.val)
            return ans

        return dfs(root, [])
