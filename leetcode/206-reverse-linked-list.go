package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func main() {
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 5}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	reversedHead := reverse(head)

	for reversedHead != nil {
		fmt.Println(reversedHead.Val)
		reversedHead = reversedHead.Next
	}
}
