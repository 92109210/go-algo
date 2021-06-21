package main

import (
	"go-algo/tree"
)

func main() {
	for _, root := range generateTrees(3) {
		tree.ShowTrees(root)
	}
}

func generateTrees(n int) []*tree.TreeNode {
	return dg(1, n)
}

func dg(left, right int) []*tree.TreeNode {
	if left > right {
		// 注意这里需要添加一个空指针元素，以便后面进入循环
		return []*tree.TreeNode{nil}
	}
	rs := make([]*tree.TreeNode, 0, 0)
	for i := left; i <= right; i++ {
		leftNode := dg(left, i-1)
		rightNode := dg(i+1, right)
		for _, l := range leftNode {
			for _, r := range rightNode {
				root := &tree.TreeNode{Val: i}
				root.Left = l
				root.Right = r
				rs = append(rs, root)
			}
		}
	}
	return rs
}
