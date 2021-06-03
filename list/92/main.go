package main

import "fmt"

func main() {
	list := []int{3, 5}
	head := getListNode(list)
	showListNode(head)
	after := reverseBetween(head, 1, 2)
	showListNode(after)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListNode(list []int) *ListNode {
	head := &ListNode{}
	pre := head
	for _, val := range list {
		temp := &ListNode{
			Val:  val,
			Next: nil,
		}
		pre.Next = temp
		pre = temp
	}
	return head.Next
}

func showListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%d, ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	node := head
	stack := make([]*ListNode, 0)
	pre := &ListNode{Next: head}
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
