from typing import List

import utils


class Solution:
    def right_side_view(self, root: utils.TreeNode) -> List[int]:
        """ 先层序遍历一遍二叉树，每层获取最右边值"""
        if not root:
            return []
        stage = [(0, root)]
        level = []
        while stage:
            count, node = stage.pop(0)
            if len(level) <= count:
                level.append([node.val])
            else:
                level[count].append(node.val)
            if node.right:
                stage.append((count + 1, node.right))
            if node.left:
                stage.append((count + 1, node.left))

        ans = [i[0] for i in level]
        return ans
