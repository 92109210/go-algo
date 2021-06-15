package main

import (
	"fmt"
	"go-algo/tree"
)

// 未完成

var slice []int

func main() {
	root := tree.GetTrees([]int{5, 3, 6, 2, 4, -1, -1, 1})
	fmt.Println(kthSmallest(root, 0))
}

func kthSmallest(root *tree.TreeNode, k int) int {
	dg(root)
	return slice[k]
}

func dg(root *tree.TreeNode) {
	if root == nil {
		return
	}
	dg(root.Left)
	slice = append(slice, root.Val)
	dg(root.Right)
}
