package main

import (
	"go-algo/tree"
)

var rs int = 0

func main() {
	root := tree.GetTrees([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8})
	result := convertBST(root)
	tree.ShowTrees(result)
}

func convertBST(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	convertBST(root.Right)
	rs += root.Val
	root.Val = rs
	convertBST(root.Left)
	return root
}
