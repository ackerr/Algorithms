from typing import List


from utils import TreeNode


class Solution:
    def level_order_bottom(self, root: TreeNode) -> List[List[int]]:
        """ use stack """
        stack = [(root, 0)]
        ans = []
        while stack:
            node, level = stack.pop()
            if node:
                if len(ans) < level + 1:
                    ans.append([])
                ans[level].append(node.val)
                stack.append((node.right, level + 1))
                stack.append((node.left, level + 1))
        return ans[::-1]
