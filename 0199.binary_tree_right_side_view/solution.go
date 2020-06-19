package leetcode0199

import "github.com/Ackerr/Algorithms/utils"

func rightSideView(root *utils.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	nodes := []*utils.TreeNode{root}
	for len(nodes) > 0{
		n := len(nodes)
		for i:=0;i<n;i++ {
			p := nodes[i]
			if p.Right != nil {
				nodes = append(nodes, p.Right)
			}
			if p.Left != nil {
				nodes = append(nodes, p.Left)
			}
			if i == 0{
				res = append(res, p.Val)
			}
		}
		nodes = nodes[n:]

	}
	return res
}
