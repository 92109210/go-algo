package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetTrees(a []int) *TreeNode {
	queue := make([]*TreeNode, 0, 0)
	tree := &TreeNode{Val: a[0]}
	head := tree
	queue = append(queue, tree)
	for i := 1; i < len(a); i++ {
		if a[i] == -1 {
			if i%2 == 0 {
				queue = queue[1:]
			}
			continue
		}
		tree = &TreeNode{Val: a[i]}
		queue = append(queue, tree)
		pre := queue[0]
		if i%2 == 1 {
			pre.Left = tree
		} else {
			pre.Right = tree
			queue = queue[1:]
		}
	}
	return head
}

// 遍历
func ShowTrees(root *TreeNode) {
	queue := make([]*TreeNode, 0, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d, ", node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Printf("\n")
}

// 前序遍历
func ShowTreesQ(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d, ", root.Val)
	ShowTreesQ(root.Left)
	ShowTreesQ(root.Right)
}
