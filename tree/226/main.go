package main

import (
	"go-algo/tree"
)

func main() {
	a := []int{4, 2, 7, 1, 3, 6, 9}
	root := tree.GetTrees(a)
	invertTree(root)
	tree.ShowTrees(root)

}

func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
