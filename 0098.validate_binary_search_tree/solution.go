package main

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    return dfs(root, -1<<63, 1<<63 - 1)
}

func dfs(node *TreeNode, low, high int) bool {
    if node == nil{
        return true
    }
    if node.Val <= low ||  node.Val >= high{
        return false
    }
    if !dfs(node.Left, low, node.Val){
        return false
    }
    if !dfs(node.Right, node.Val, high){
        return false
    }
    return true
}
