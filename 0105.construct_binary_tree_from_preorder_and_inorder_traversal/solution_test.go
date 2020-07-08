package leetcode0105

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	tree := buildTree(preorder, inorder)
	if reflect.DeepEqual(buildInorder(tree, []int{}), inorder){
		t.Errorf("buildTree(%+v,  %+v), result error ", preorder, inorder)
	}

}

func buildInorder(node *utils.TreeNode, inorder []int) []int {
	if node == nil{
		return inorder
	}
	if node.Left != nil {
		buildInorder(node.Left, inorder)
	}
	inorder = append(inorder, node.Val)
	if node.Right != nil {
		buildInorder(node.Right, inorder)
	}
	return inorder
}
