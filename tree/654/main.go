package main

import (
	"go-algo/tree"
)

func main() {
}

func constructMaximumBinaryTree(nums []int) *tree.TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	index := getMaxValueIndex(nums)
	node := tree.TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[0:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
	return &node
}

func getMaxValueIndex(nums []int) (index int) {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	for i, a := range nums {
		if a > nums[index] {
			index = i
		}
	}
	return index
}
