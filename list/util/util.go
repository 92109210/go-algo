package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetListNode(list []int) *ListNode {
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

func ShowListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%d, ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
