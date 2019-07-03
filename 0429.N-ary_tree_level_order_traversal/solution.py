class Node:
    def __init__(self, val, children):
        self.val = val
        self.children = children


class Solution:
    def level_order(self, root: "Node") -> List[List[int]]:
        ans, nodes = [], [root]
        while nodes:
            temp, temp_nodes = [], []
            for node in nodes:
                temp.append(node.val)
                temp_nodes.extend(node.children)
            ans.append(temp)
            nodes = temp_nodes
        return ans
