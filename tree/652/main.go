package main

import (
	"fmt"
	"go-algo/tree"
	"strconv"
)

var RS = make(map[string]int,10)

func main() {
	root := tree.GetTrees([]int{0, 0, 0, 0, -1, -1, 0, -1, -1, 0, 0})
	//tree.ShowTrees(root)
	rs := findDuplicateSubtrees(root)
	for _, r := range rs {
		tree.ShowTrees(r)
	}
}

func findDuplicateSubtrees(root *tree.TreeNode) []*tree.TreeNode {
	if root == nil {
		return nil
	}
	t := make([]*tree.TreeNode, 0)
	t = append(t, findDuplicateSubtrees(root.Left)...)
	t = append(t, findDuplicateSubtrees(root.Right)...)
	str := getrs(root)
	fmt.Println(str)
	RS[str] = RS[str] + 1
	if RS[str] == 2 {
		fmt.Printf("被添加入结果的str：%s\n", str)
		t = append(t, root)
	}
	return t
}

func getrs(root *tree.TreeNode) string {
	s := ""
	if root == nil {
		s = "#"
		return s
	}
	s = strconv.Itoa(root.Val) + ", " + getrs(root.Left) + ", " + getrs(root.Right)
	return s
}
