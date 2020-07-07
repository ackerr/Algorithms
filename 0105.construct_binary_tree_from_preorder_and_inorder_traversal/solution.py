import utils
from typing import List


class Solution:
    def build_tree(self, preorder: List[int], inorder: List[int]) -> utils.TreeNode:
        """
        前序遍历： [根节点，[左子树前序遍历]， [右子树前序遍历]]
        中序遍历： [[左子树中序遍历], 根节点, [右子树中序遍历]]
        """

        def _build(preorder_left, preorder_right, inorder_left, inorder_right):
            if preorder_left > preorder_right:
                return None

            root = utils.TreeNode(preorder[preorder_left])  # 根节点
            inorder_root = index[preorder[preorder_left]]  # 根节点在中序遍历中的索引
            left_length = inorder_root - inorder_left  # 左子树的长度
            root.left = _build(preorder_left + 1, preorder_left + left_length, inorder_left, inorder_root - 1)
            root.right = _build(preorder_left + left_length + 1, preorder_right, inorder_root + 1, inorder_right)
            return root

        n = len(preorder) - 1
        index = {val: i for i, val in enumerate(inorder)}
        return _build(0, n, 0, n)
