package leetcode0102

import "github.com/Ackerr/Algorithms/utils"

func levelOrder(root *utils.TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		temp := make([]int, length)
		for i := 0; i < length; i++ {
			set := queue[i]
			temp[i] = set.Val
			if set.Left != nil {
				queue = append(queue, set.Left)
			}
			if set.Right != nil {
				queue = append(queue, set.Right)
			}
		}
		ans = append(ans, temp)
		queue = queue[length:]
	}
	return ans
}
