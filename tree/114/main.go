package main

import (
	"go-algo/tree"
)

func main() {
	t := tree.GetTrees([]int{1, 2, 5, 3, 4, -1, 6})
	tree.ShowTrees(t)
	flatten(t)
	tree.ShowTrees(t)
}

func flatten(root *tree.TreeNode) {

	if root == nil {
		return
	}

	// 前序遍历位置
	flatten(root.Left)
	// 中序遍历位置
	flatten(root.Right)

	// 后续遍历位置
	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}
