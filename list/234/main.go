package main

import (
	"fmt"
	"go-algo/list/util"
)

// 使用快慢指针找到中间的位置，
// 从中间位置翻转后面的链表，
// 比较连个链表，
// 再次翻转后面的链表恢复成原来的链表

func main() {
	a := util.GetListNode([]int{1, 2, 3, 4, 4, 2, 1})
	fmt.Println(isPalindrome(a))
	util.ShowListNode(a)

}

func isPalindrome(head *util.ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 反转slow开始的链表
	head2 := fz(slow)
	rec := head2
	for head2 != nil {
		if head2.Val != head.Val {
			// 恢复链表
			fz(rec)
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	// 恢复链表
	fz(rec)
	return true

}

func fz(head *util.ListNode) *util.ListNode {
	var pre *util.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
