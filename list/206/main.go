package main

import (
	"go-algo/list/util"
)

func main() {
	head := util.GetListNode([]int{1, 2, 3, 4, 5})
	util.ShowListNode(head)
	rs := dd(head)
	util.ShowListNode(rs)
}

// 递归解法
func dg(node *util.ListNode) *util.ListNode {
	if node.Next == nil {
		return node
	}
	rs := dg(node.Next)
	node.Next.Next = node
	node.Next = nil
	return rs
}

// 迭代
func dd(node *util.ListNode) *util.ListNode {
	curr := node
	var pre *util.ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
