package leetcode0105

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)

	if !reflect.DeepEqual(buildInorder(tree), inorder) {
		t.Errorf("buildTree(%+v,  %+v), result error ", preorder, inorder)
	}

}

func buildInorder(node *utils.TreeNode) []int {
	return dfs(node)
}

func dfs(node *utils.TreeNode) []int {
	if node == nil {
		return nil
	}
	var ans []int
	ans = append(ans, dfs(node.Left)...)
	ans = append(ans, node.Val)
	ans = append(ans, dfs(node.Right)...)
	return ans
}
