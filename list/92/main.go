package main

import (
	"go-algo/list/util"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	head := util.GetListNode(list)
	util.ShowListNode(head)
	after := reverseBetween1(head, 2, 4)
	util.ShowListNode(after)
}

func reverseBetween(head *util.ListNode, left int, right int) *util.ListNode {
	node := head
	stack := make([]*util.ListNode, 0)
	pre := &util.ListNode{Next: head}
	prehead := pre
	index := 0
	listIndex := 1
	for node != nil {
		if listIndex >= left && listIndex <= right {
			stack = append(stack, node)
			index++
			node = node.Next
			listIndex++
			continue
		}
		if index != 0 {
			break
		}
		listIndex++
		pre = node
		node = node.Next
	}
	for index != 0 {
		pre.Next = stack[index-1]
		pre = pre.Next
		index--
	}
	pre.Next = node
	return prehead.Next
}

// 优化
// 头插法
func reverseBetween1(head *util.ListNode, left int, right int) *util.ListNode {
	pre := &util.ListNode{Next: head}
	head = pre
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	curr := pre.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return head.Next
}
