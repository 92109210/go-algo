package main

import "go-algo/tree"

func main() {
	root := tree.GetTrees([]int{4, 2, 7, 1, 3})
	rs := searchBST(root, 2)
	tree.ShowTrees(rs)

}

func searchBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := searchBST(root.Left, val)
	right := searchBST(root.Right, val)
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
