class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def binary_tree_paths(self, root: TreeNode) -> List[str]:
        paths = []
        
        def deep(root, path):
            if not root:
                return
            path = path + [str(root.val)]
            if not root.left and not root.right:
                return paths.append('->'.join(path))
            deep(root.left, path)
            deep(root.right, path)
            
        deep(root, [])
        return paths