package main

import "go-algo/tree"

func main() {

}

func buildTree(inorder []int, postorder []int) *tree.TreeNode {
	if postorder == nil || len(postorder) == 0 {
		return nil
	}
	index := getIndexByValue(inorder, postorder[len(postorder)-1])
	lp := postorder[:index]
	rp := postorder[index : len(postorder)-1]
	li := inorder[:index]
	ri := inorder[index+1:]
	node := tree.TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(li, lp),
		Right: buildTree(ri, rp),
	}
	return &node
}

func getIndexByValue(nums []int, val int) (index int) {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return index
}
