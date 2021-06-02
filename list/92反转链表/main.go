package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	node := head
	stack := make([]*ListNode, 0)
	index := 0
	listIndex := 1
	for node != nil {
		if listIndex >= left && listIndex < right {
			stack = append(stack, node.Next)
			index++
			node = node.Next
			break
		}
		if index != 0 {
			temp := node
			node = node.Next
			temp.Next = stack[index-1]

		}
	}
}
