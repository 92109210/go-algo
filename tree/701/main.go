package main

import "go-algo/tree"

func main() {

}

func insertIntoBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		temp := &tree.TreeNode{Val: val}
		return temp
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
