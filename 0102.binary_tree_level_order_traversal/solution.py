from typing import List


from utils import TreeNode


class Solution:
    def level_order(self, root: TreeNode) -> List[List[int]]:
        """ 广度优先 """
        if not root:
            return []
        queue = [(root, 0)]
        ans = []
        while queue:
            node, level = queue.pop(0)
            if level > len(ans) - 1:
                ans.append([node.val])
            else:
                ans[level].append(node.val)
            if node.left:
                queue.append((node.left, level + 1))
            if node.right:
                queue.append((node.right, level + 1))
        return ans
