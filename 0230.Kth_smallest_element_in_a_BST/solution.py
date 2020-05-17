from utils import TreeNode


class Solution:
    def kth_smallest(self, root: TreeNode, k: int) -> int:
        def middle(node):  # 二叉搜索树中序遍历得到的是递增数组
            if not node:
                return []
            return middle(node.left) + [node.val] + middle(node.right)

        return middle(root)[k - 1]
