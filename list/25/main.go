package main

import (
	"go-algo/list/util"
)

func main() {
	head := util.GetListNode([]int{1, 2, 3, 4, 5})
	util.ShowListNode(head)
	rs := reverseKGroup1(head, 2)
	util.ShowListNode(rs)
}

// 递归
func reverseKGroup(head *util.ListNode, k int) *util.ListNode {
	j := k
	curr := head
	for j != 1 && curr != nil {
		curr = curr.Next
		j--
	}
	if curr == nil {
		return head
	}
	next := curr.Next
	curr.Next = nil
	rs := dg(head)
	head.Next = reverseKGroup(next, k)
	return rs
}

// 递归
func dg(node *util.ListNode) *util.ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	rs := dg(node.Next)
	node.Next.Next = node
	node.Next = nil
	return rs
}

// 迭代
func reverseKGroup1(head *util.ListNode, k int) *util.ListNode {
	rs := &util.ListNode{}
	p1 := head
	p2 := head
	i := 0
	flag := 1
	var next *util.ListNode
	var last *util.ListNode
	for p2 != nil {
		i++
		if i == k {
			next = p2.Next
			p2.Next = nil
			if flag == 1 {
				rs = dd(p1)
				flag = 0
			} else {
				last.Next = dd(p1)
			}
			p1.Next = next
			last = p1
			p2 = p1
			p1 = next
			i = 0
		}
		p2 = p2.Next
	}
	return rs
}

// 迭代
func dd(node *util.ListNode) *util.ListNode {
	var pre *util.ListNode
	curr := node
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
