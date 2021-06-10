package main

import "go-algo/tree"

func main() {

}

func t() *tree.TreeNode {
	return nil
}

func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	inorderIndex := getIndexByValue(inorder, preorder[0])
	leftPerprder := preorder[1 : inorderIndex+1]
	rightPerprder := preorder[inorderIndex+1:]
	leftInorder := inorder[:inorderIndex]
	rightInorder := inorder[inorderIndex+1:]
	node := tree.TreeNode{
		Val:   preorder[0],
		Left:  buildTree(leftPerprder, leftInorder),
		Right: buildTree(rightPerprder, rightInorder),
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
