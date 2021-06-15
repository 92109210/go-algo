package main

import (
	"go-algo/tree"
)

func main() {
	root := tree.GetTrees([]int{5, 3, 6, 2, 4, -1, 7})
	rs := deleteNode(root, 3)
	tree.ShowTrees(rs)
}

func deleteNode(root *tree.TreeNode, key int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 找到了
		temp := root.Left
		if temp == nil {
			root = root.Right
			return root
		}
		for temp.Right != nil {
			temp = temp.Right
		}
		temp.Right = root.Right
		root = root.Left
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
