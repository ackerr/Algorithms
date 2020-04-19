class Solution:
    def path_sum(self, root, sum: int) -> int:
        """ 遍历所有的路径以及子路径，每次递归就将(sum - val), 出现0是，说明找到符合条件的 """
        if not root:
            return 0

        def deep(root, sum):
            count = 0
            if not root:
                return count
            if root.val == sum:
                count += 1
            count += deep(root.left, sum - root.val)
            count += deep(root.right, sum - root.val)
            return count

        return deep(root, sum) + self.path_sum(root.left, sum) + self.path_sum(root.right, sum)
