from utils import TreeNode


class Solution:
    def is_valid_bst(self, root: TreeNode) -> bool:
        """ 判断中序遍历后是否是升序数组 """
        if not root:
            return True
        ans = []

        def in_order(node):
            if not node:
                return
            in_order(node.left)
            ans.append(node.val)
            in_order(node.right)

        in_order(root)
        for i in range(1, len(ans)):
            if ans[i - 1] >= ans[i]:
                return False
        return True

    def recur(self, root: TreeNode) -> bool:
        """ 递归判断每个节点左节点都小于该节点，右节点都大于该节点 """

        def dfs(node, left, right):
            if not node:
                return True
            if node.val < left or node.val > right:
                return False
            if not dfs(node.left, left, node.val):
                return False
            if not dfs(node.right, node.val, right):
                return False
            return True

        return dfs(root, float("-inf"), float("inf"))
