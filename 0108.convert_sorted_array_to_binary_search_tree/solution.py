class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def sorted_array_to_BST(self, nums: List[int]) -> TreeNode:
        if not nums:
            return
        # 通过中序遍历可知，根节点所在数组位置的中间
        mid = len(nums) // 2
        root = TreeNode(mid)
        root.left = self.sorted_array_to_BST(nums[0:mid])
        root.right = self.sorted_array_to_BST(nums[mid + 1 : len(nums)])
        return root
