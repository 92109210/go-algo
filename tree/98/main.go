package main

import (
	"fmt"
	"go-algo/tree"
	"math"
)

func main() {
	root := tree.GetTrees([]int{5,1,4,-1,-1,3,6})
	fmt.Println(isValidBST(root))

}

func isValidBST(root *tree.TreeNode) bool {
	return dg(root, math.MinInt64, math.MaxInt64)
}

func dg(root *tree.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	rs1 := true
	if root.Left != nil {
		rs1 = rs1 && (root.Val > root.Left.Val)
		rs1 = rs1 && (root.Left.Val > min)
	}
	if root.Right != nil {
		rs1 = rs1 && (root.Val < root.Right.Val)
		rs1 = rs1 && (root.Right.Val < max)
	}
	if rs1 {
		rs1 = rs1 && dg(root.Left, min, root.Val)
		rs1 = rs1 && dg(root.Right, root.Val, max)
	}
	return rs1
}
