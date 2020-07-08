package leetcode0105

import "github.com/Ackerr/Algorithms/utils"

func buildTree(preorder []int, inorder []int) *utils.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &utils.TreeNode{Val: preorder[0]}
	index := 0
	for i, val := range inorder {
		if val == preorder[0] {
			index = i
			break
		}
	}
	leftLength := len(inorder[:index])
	node.Left = buildTree(preorder[1:leftLength+1], inorder[:index])
	node.Right = buildTree(preorder[leftLength+1:], inorder[index+1:])
	return node
}
